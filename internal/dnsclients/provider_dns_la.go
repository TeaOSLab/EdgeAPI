package dnsclients

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	teaconst "github.com/TeaOSLab/EdgeAPI/internal/const"
	"github.com/TeaOSLab/EdgeAPI/internal/dnsclients/dnsla"
	"github.com/TeaOSLab/EdgeAPI/internal/dnsclients/dnstypes"
	"github.com/TeaOSLab/EdgeAPI/internal/errors"
	"github.com/iwind/TeaGo/maps"
	"github.com/iwind/TeaGo/types"
)

const DNSLaAPIEndpoint = "https://api.dns.la"

var dnsLAHTTPClient = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

type DNSLaProvider struct {
	BaseProvider

	ProviderId int64

	apiId  string
	secret string

	routesLocker sync.Mutex
	cachedRoutes map[string][]*dnstypes.Route // domain => []Route
}

// Auth 认证
func (this *DNSLaProvider) Auth(params maps.Map) error {
	this.apiId = params.GetString("apiId")
	this.secret = params.GetString("secret")

	if len(this.apiId) == 0 {
		return errors.New("'apiId' should not be empty")
	}
	if len(this.secret) == 0 {
		return errors.New("'secret' should not be empty")
	}

	this.cachedRoutes = map[string][]*dnstypes.Route{}

	return nil
}

// MaskParams 对参数进行掩码
func (this *DNSLaProvider) MaskParams(params maps.Map) {
	if params == nil {
		return
	}
	params["secret"] = MaskString(params.GetString("secret"))
}

// GetDomains 获取所有域名列表
func (this *DNSLaProvider) GetDomains() (domains []string, err error) {
	for i := 1; i < 5000; i++ {
		var resp = &dnsla.DomainListResponse{}
		err = this.doAPI(http.MethodGet, "/api/domainList", map[string]string{
			"pageSize":  "100",
			"pageIndex": types.String(i),
		}, nil, resp)
		if err != nil {
			return nil, err
		}
		if !resp.Success() {
			return nil, resp.Error()
		}

		if len(resp.Data.Results) == 0 {
			return
		}

		for _, data := range resp.Data.Results {
			domains = append(domains, strings.TrimSuffix(data.Domain, "."))
		}
	}
	return
}

// GetRecords 获取域名解析记录列表
func (this *DNSLaProvider) GetRecords(domain string) (records []*dnstypes.Record, err error) {
	domainId, err := this.getDomainId(domain)
	if err != nil {
		return nil, err
	}
	if len(domainId) == 0 {
		return
	}

	for i := 1; i < 5000; i++ {
		var resp = &dnsla.RecordListResponse{}
		err = this.doAPI(http.MethodGet, "/api/recordList", map[string]string{
			"domainId":  domainId,
			"pageSize":  "100",
			"pageIndex": types.String(i),
		}, nil, resp)
		if err != nil {
			return
		}
		if !resp.Success() {
			return nil, resp.Error()
		}
		if len(resp.Data.Results) == 0 {
			break
		}
		for _, rawRecord := range resp.Data.Results {
			var recordType = this.recordTypeName(rawRecord.Type)

			// 修正Record
			if recordType == dnstypes.RecordTypeCNAME && !strings.HasSuffix(rawRecord.Data, ".") {
				rawRecord.Data += "."
			}

			records = append(records, &dnstypes.Record{
				Id:    rawRecord.Id,
				Name:  rawRecord.Host,
				Type:  recordType,
				Value: rawRecord.Data,
				Route: rawRecord.LineCode,
				TTL:   types.Int32(rawRecord.TTL),
			})
		}
	}

	// 写入缓存
	if this.ProviderId > 0 {
		sharedDomainRecordsCache.WriteDomainRecords(this.ProviderId, domain, records)
	}

	return
}

// GetRoutes 读取域名支持的线路数据
func (this *DNSLaProvider) GetRoutes(domain string) (routes []*dnstypes.Route, err error) {
	var resp = &dnsla.AllLineListResponse{}
	err = this.doAPI(http.MethodGet, "/api/allLineList", nil, nil, resp)
	if err != nil {
		return
	}
	if !resp.Success() {
		return nil, resp.Error()
	}

	for _, data := range resp.Data {
		routes = append(routes, &dnstypes.Route{
			Name: data.Name,
			Code: data.Id + "$" + data.Code, // ID + $ + CODE
		})
		routes = append(routes, this.travelLines(data.Children)...)
	}

	this.routesLocker.Lock()
	this.cachedRoutes[domain] = routes
	this.routesLocker.Unlock()

	return
}

// QueryRecord 查询单个记录
func (this *DNSLaProvider) QueryRecord(domain string, name string, recordType dnstypes.RecordType) (*dnstypes.Record, error) {
	// 从缓存中读取
	if this.ProviderId > 0 {
		record, hasRecords, _ := sharedDomainRecordsCache.QueryDomainRecord(this.ProviderId, domain, name, recordType)
		if hasRecords { // 有效的搜索
			return record, nil
		}
	}

	domainId, err := this.getDomainId(domain)
	if err != nil {
		return nil, err
	}
	if len(domainId) == 0 {
		return nil, nil
	}

	var resp = &dnsla.RecordListResponse{}
	err = this.doAPI(http.MethodGet, "/api/recordList", map[string]string{
		"domainId":  domainId,
		"pageSize":  "100",
		"pageIndex": "1",
		"host":      name,
		"type":      types.String(this.recordTypeId(recordType)),
	}, nil, resp)
	if err != nil {
		return nil, err
	}
	if !resp.Success() {
		return nil, resp.Error()
	}
	if len(resp.Data.Results) == 0 {
		return nil, nil
	}
	for _, rawRecord := range resp.Data.Results {
		var recordTypeName = this.recordTypeName(rawRecord.Type)

		if rawRecord.Host == name && recordTypeName == recordType {
			// 修正Record
			if recordType == dnstypes.RecordTypeCNAME && !strings.HasSuffix(rawRecord.Data, ".") {
				rawRecord.Data += "."
			}

			return &dnstypes.Record{
				Id:    rawRecord.Id,
				Name:  rawRecord.Host,
				Type:  recordTypeName,
				Value: rawRecord.Data,
				Route: rawRecord.LineCode,
				TTL:   types.Int32(rawRecord.TTL),
			}, nil
		}
	}

	return nil, nil
}

// QueryRecords 查询多个记录
func (this *DNSLaProvider) QueryRecords(domain string, name string, recordType dnstypes.RecordType) ([]*dnstypes.Record, error) {
	// 从缓存中读取
	if this.ProviderId > 0 {
		records, hasRecords, _ := sharedDomainRecordsCache.QueryDomainRecords(this.ProviderId, domain, name, recordType)
		if hasRecords { // 有效的搜索
			return records, nil
		}
	}

	domainId, err := this.getDomainId(domain)
	if err != nil {
		return nil, err
	}
	if len(domainId) == 0 {
		return nil, nil
	}

	var result []*dnstypes.Record
	for pageIndex := 1; pageIndex < 5000; pageIndex++ {
		var resp = &dnsla.RecordListResponse{}
		err = this.doAPI(http.MethodGet, "/api/recordList", map[string]string{
			"domainId":  domainId,
			"pageSize":  "100",
			"pageIndex": types.String(pageIndex),
			"host":      name,
			"type":      types.String(this.recordTypeId(recordType)),
		}, nil, resp)
		if err != nil {
			return nil, err
		}
		if !resp.Success() {
			return nil, resp.Error()
		}
		if len(resp.Data.Results) == 0 {
			break
		}
		for _, rawRecord := range resp.Data.Results {
			var recordTypeName = this.recordTypeName(rawRecord.Type)
			if rawRecord.Host == name && recordTypeName == recordType {

				// 修正Record
				if recordType == dnstypes.RecordTypeCNAME && !strings.HasSuffix(rawRecord.Data, ".") {
					rawRecord.Data += "."
				}

				result = append(result, &dnstypes.Record{
					Id:    rawRecord.Id,
					Name:  rawRecord.Host,
					Type:  recordTypeName,
					Value: rawRecord.Data,
					Route: rawRecord.LineCode,
					TTL:   types.Int32(rawRecord.TTL),
				})
			}
		}
	}

	return result, nil
}

// AddRecord 设置记录
func (this *DNSLaProvider) AddRecord(domain string, newRecord *dnstypes.Record) error {
	routeId, err := this.routeToId(domain, newRecord.Route)
	if err != nil {
		return err
	}

	var ttl = newRecord.TTL
	if ttl <= 0 {
		ttl = 600
	}

	domainId, err := this.getDomainId(domain)
	if err != nil {
		return err
	}

	if newRecord.Type == dnstypes.RecordTypeCNAME && !strings.HasSuffix(newRecord.Value, ".") {
		newRecord.Value += "."
	}

	recordJSON, err := json.Marshal(map[string]any{
		"domainId": domainId,
		"host":     newRecord.Name,
		"type":     this.recordTypeId(newRecord.Type),
		"data":     newRecord.Value,
		"ttl":      ttl,
		"lineId":   routeId,
	})
	if err != nil {
		return err
	}

	var resp = &dnsla.RecordCreateResponse{}
	err = this.doAPI(http.MethodPost, "/api/record", nil, recordJSON, resp)
	if err != nil {
		return err
	}
	if !resp.Success() {
		return resp.Error()
	}
	newRecord.Id = types.String(resp.Data.Id)

	// 加入缓存
	if this.ProviderId > 0 {
		sharedDomainRecordsCache.AddDomainRecord(this.ProviderId, domain, newRecord)
	}

	return nil
}

// UpdateRecord 修改记录
func (this *DNSLaProvider) UpdateRecord(domain string, record *dnstypes.Record, newRecord *dnstypes.Record) error {
	if len(record.Id) == 0 {
		return errors.New("record id required")
	}

	routeId, err := this.routeToId(domain, newRecord.Route)
	if err != nil {
		return err
	}

	var ttl = newRecord.TTL
	if ttl <= 0 {
		ttl = 600
	}

	domainId, err := this.getDomainId(domain)
	if err != nil {
		return err
	}

	if newRecord.Type == dnstypes.RecordTypeCNAME && !strings.HasSuffix(newRecord.Value, ".") {
		newRecord.Value += "."
	}

	recordJSON, err := json.Marshal(map[string]any{
		"id":       record.Id,
		"domainId": domainId,
		"host":     newRecord.Name,
		"type":     this.recordTypeId(newRecord.Type),
		"data":     newRecord.Value,
		"ttl":      ttl,
		"lineId":   routeId,
	})
	if err != nil {
		return err
	}

	var resp = &dnsla.RecordUpdateResponse{}
	err = this.doAPI(http.MethodPut, "/api/record", nil, recordJSON, resp)
	if err != nil {
		return err
	}
	if !resp.Success() {
		return resp.Error()
	}
	newRecord.Id = record.Id

	// 修改缓存
	if this.ProviderId > 0 {
		sharedDomainRecordsCache.UpdateDomainRecord(this.ProviderId, domain, newRecord)
	}

	return nil
}

// DeleteRecord 删除记录
func (this *DNSLaProvider) DeleteRecord(domain string, record *dnstypes.Record) error {
	var resp = &dnsla.RecordDeleteResponse{}
	err := this.doAPI(http.MethodDelete, "/api/record", map[string]string{
		"id": record.Id,
	}, nil, resp)
	if err != nil {
		return err
	}
	if !resp.Success() {
		// ignore not found error
		if resp.Code == 404 {
			return nil
		}

		return resp.Error()
	}

	// 删除缓存
	if this.ProviderId > 0 {
		sharedDomainRecordsCache.DeleteDomainRecord(this.ProviderId, domain, record.Id)
	}

	return nil
}

// DefaultRoute 默认线路
func (this *DNSLaProvider) DefaultRoute() string {
	return "default"
}

// 发送请求
func (this *DNSLaProvider) doAPI(method string, path string, params map[string]string, postJSONData []byte, respPtr interface{}) error {
	var apiURL = DNSLaAPIEndpoint + path

	if len(params) > 0 {
		var query = &url.Values{}
		for k, v := range params {
			query.Set(k, v)
		}
		apiURL += "?" + query.Encode()
	}

	var bodyReader io.Reader
	if len(postJSONData) > 0 {
		bodyReader = bytes.NewReader(postJSONData)
	}

	req, err := http.NewRequest(method, apiURL, bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", teaconst.ProductName+"/"+teaconst.Version)
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(this.apiId+":"+this.secret)))

	if len(postJSONData) > 0 {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	resp, err := dnsLAHTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == 0 {
		return errors.New("invalid response status '" + strconv.Itoa(resp.StatusCode) + "', response '" + string(data) + "'")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("response error: " + string(data))
	}

	if respPtr != nil {
		err = json.Unmarshal(data, respPtr)
		if err != nil {
			return fmt.Errorf("decode json failed: %w: %s", err, string(data))
		}
	}

	return nil
}

func (this *DNSLaProvider) getDomainId(domain string) (string, error) {
	var resp = &dnsla.DomainResponse{}
	err := this.doAPI(http.MethodGet, "/api/domain", map[string]string{
		"domain": domain,
	}, nil, resp)
	if err != nil {
		return "", err
	}
	return resp.Data.Id, nil
}

func (this *DNSLaProvider) recordTypeName(recordTypeId int) string {
	switch recordTypeId {
	case 1:
		return "A"
	case 2:
		return "NS"
	case 5:
		return "CNAME"
	case 15:
		return "MX"
	case 16:
		return "TXT"
	case 28:
		return "AAAA"
	case 33:
		return "SRV"
	case 257:
		return "CAA"
	case 256:
		return "URL转发"
	}
	return "UNKNOWN"
}

func (this *DNSLaProvider) recordTypeId(recordTypeName string) int {
	switch recordTypeName {
	case "A":
		return 1
	case "NS":
		return 2
	case "CNAME":
		return 5
	case "MX":
		return 15
	case "TXT":
		return 16
	case "AAAA":
		return 28
	case "SRV":
		return 33
	case "CAA":
		return 257
	case "URL转发":
		return 256
	}
	return 0
}

func (this *DNSLaProvider) travelLines(children []dnsla.AllLineListResponseChild) (result []*dnstypes.Route) {
	if len(children) == 0 {
		return
	}
	for _, child := range children {
		result = append(result, &dnstypes.Route{
			Name: child.Name,
			Code: child.Id + "$" + child.Code,
		})
		result = append(result, this.travelLines(child.Children)...)
	}
	return
}

func (this *DNSLaProvider) routeToId(domain string, routeCode string) (string, error) {
	if len(routeCode) == 0 {
		return "", nil
	}
	if routeCode == "default" {
		return "", nil
	}

	// 新的线路：id@code
	if strings.Contains(routeCode, "$") {
		return strings.Split(routeCode, "$")[0], nil
	}

	// 兼容老的线路
	this.routesLocker.Lock()
	var hasCachedRoutes = len(this.cachedRoutes[domain]) > 0
	this.routesLocker.Unlock()

	if !hasCachedRoutes {
		_, err := this.GetRoutes(domain)
		if err != nil {
			return "", err
		}
	}

	this.routesLocker.Lock()
	defer this.routesLocker.Unlock()
	if len(this.cachedRoutes) == 0 {
		return "", nil
	}

	for _, cachedRoute := range this.cachedRoutes[domain] {
		if strings.HasSuffix(cachedRoute.Code, "$"+routeCode) {
			return strings.Split(cachedRoute.Code, "$")[0], nil
		}
	}

	return "", errors.New("invalid route code '" + routeCode + "'")
}
