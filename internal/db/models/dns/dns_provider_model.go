package dns

import "github.com/iwind/TeaGo/dbs"

const (
	DNSProviderField_Id            dbs.FieldName = "id"            // ID
	DNSProviderField_Name          dbs.FieldName = "name"          // 名称
	DNSProviderField_AdminId       dbs.FieldName = "adminId"       // 管理员ID
	DNSProviderField_UserId        dbs.FieldName = "userId"        // 用户ID
	DNSProviderField_Type          dbs.FieldName = "type"          // 供应商类型
	DNSProviderField_ApiParams     dbs.FieldName = "apiParams"     // API参数
	DNSProviderField_CreatedAt     dbs.FieldName = "createdAt"     // 创建时间
	DNSProviderField_State         dbs.FieldName = "state"         // 状态
	DNSProviderField_DataUpdatedAt dbs.FieldName = "dataUpdatedAt" // 数据同步时间
	DNSProviderField_MinTTL        dbs.FieldName = "minTTL"        // 最小TTL
)

// DNSProvider DNS服务商
type DNSProvider struct {
	Id            uint32   `field:"id"`            // ID
	Name          string   `field:"name"`          // 名称
	AdminId       uint32   `field:"adminId"`       // 管理员ID
	UserId        uint32   `field:"userId"`        // 用户ID
	Type          string   `field:"type"`          // 供应商类型
	ApiParams     dbs.JSON `field:"apiParams"`     // API参数
	CreatedAt     uint64   `field:"createdAt"`     // 创建时间
	State         uint8    `field:"state"`         // 状态
	DataUpdatedAt uint64   `field:"dataUpdatedAt"` // 数据同步时间
	MinTTL        uint32   `field:"minTTL"`        // 最小TTL
}

type DNSProviderOperator struct {
	Id            any // ID
	Name          any // 名称
	AdminId       any // 管理员ID
	UserId        any // 用户ID
	Type          any // 供应商类型
	ApiParams     any // API参数
	CreatedAt     any // 创建时间
	State         any // 状态
	DataUpdatedAt any // 数据同步时间
	MinTTL        any // 最小TTL
}

func NewDNSProviderOperator() *DNSProviderOperator {
	return &DNSProviderOperator{}
}
