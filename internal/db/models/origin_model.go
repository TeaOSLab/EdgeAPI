package models

import "github.com/iwind/TeaGo/dbs"

const (
	OriginField_Id                 dbs.FieldName = "id"                 // ID
	OriginField_AdminId            dbs.FieldName = "adminId"            // 管理员ID
	OriginField_UserId             dbs.FieldName = "userId"             // 用户ID
	OriginField_ReverseProxyId     dbs.FieldName = "reverseProxyId"     // 所属反向代理ID
	OriginField_IsOn               dbs.FieldName = "isOn"               // 是否启用
	OriginField_Name               dbs.FieldName = "name"               // 名称
	OriginField_Version            dbs.FieldName = "version"            // 版本
	OriginField_Addr               dbs.FieldName = "addr"               // 地址
	OriginField_Oss                dbs.FieldName = "oss"                // OSS配置
	OriginField_Description        dbs.FieldName = "description"        // 描述
	OriginField_Code               dbs.FieldName = "code"               // 代号
	OriginField_Weight             dbs.FieldName = "weight"             // 权重
	OriginField_ConnTimeout        dbs.FieldName = "connTimeout"        // 连接超时
	OriginField_ReadTimeout        dbs.FieldName = "readTimeout"        // 读超时
	OriginField_IdleTimeout        dbs.FieldName = "idleTimeout"        // 空闲连接超时
	OriginField_MaxFails           dbs.FieldName = "maxFails"           // 最多失败次数
	OriginField_MaxConns           dbs.FieldName = "maxConns"           // 最大并发连接数
	OriginField_MaxIdleConns       dbs.FieldName = "maxIdleConns"       // 最多空闲连接数
	OriginField_HttpRequestURI     dbs.FieldName = "httpRequestURI"     // 转发后的请求URI
	OriginField_HttpRequestHeader  dbs.FieldName = "httpRequestHeader"  // 请求Header配置
	OriginField_HttpResponseHeader dbs.FieldName = "httpResponseHeader" // 响应Header配置
	OriginField_Host               dbs.FieldName = "host"               // 自定义主机名
	OriginField_HealthCheck        dbs.FieldName = "healthCheck"        // 健康检查设置
	OriginField_Cert               dbs.FieldName = "cert"               // 证书设置
	OriginField_Ftp                dbs.FieldName = "ftp"                // FTP相关设置
	OriginField_CreatedAt          dbs.FieldName = "createdAt"          // 创建时间
	OriginField_Domains            dbs.FieldName = "domains"            // 所属域名
	OriginField_FollowPort         dbs.FieldName = "followPort"         // 端口跟随
	OriginField_State              dbs.FieldName = "state"              // 状态
	OriginField_Http2Enabled       dbs.FieldName = "http2Enabled"       // 是否支持HTTP/2
)

// Origin 源站
type Origin struct {
	Id                 uint32   `field:"id"`                 // ID
	AdminId            uint32   `field:"adminId"`            // 管理员ID
	UserId             uint32   `field:"userId"`             // 用户ID
	ReverseProxyId     uint64   `field:"reverseProxyId"`     // 所属反向代理ID
	IsOn               bool     `field:"isOn"`               // 是否启用
	Name               string   `field:"name"`               // 名称
	Version            uint32   `field:"version"`            // 版本
	Addr               dbs.JSON `field:"addr"`               // 地址
	Oss                dbs.JSON `field:"oss"`                // OSS配置
	Description        string   `field:"description"`        // 描述
	Code               string   `field:"code"`               // 代号
	Weight             uint32   `field:"weight"`             // 权重
	ConnTimeout        dbs.JSON `field:"connTimeout"`        // 连接超时
	ReadTimeout        dbs.JSON `field:"readTimeout"`        // 读超时
	IdleTimeout        dbs.JSON `field:"idleTimeout"`        // 空闲连接超时
	MaxFails           uint32   `field:"maxFails"`           // 最多失败次数
	MaxConns           uint32   `field:"maxConns"`           // 最大并发连接数
	MaxIdleConns       uint32   `field:"maxIdleConns"`       // 最多空闲连接数
	HttpRequestURI     string   `field:"httpRequestURI"`     // 转发后的请求URI
	HttpRequestHeader  dbs.JSON `field:"httpRequestHeader"`  // 请求Header配置
	HttpResponseHeader dbs.JSON `field:"httpResponseHeader"` // 响应Header配置
	Host               string   `field:"host"`               // 自定义主机名
	HealthCheck        dbs.JSON `field:"healthCheck"`        // 健康检查设置
	Cert               dbs.JSON `field:"cert"`               // 证书设置
	Ftp                dbs.JSON `field:"ftp"`                // FTP相关设置
	CreatedAt          uint64   `field:"createdAt"`          // 创建时间
	Domains            dbs.JSON `field:"domains"`            // 所属域名
	FollowPort         bool     `field:"followPort"`         // 端口跟随
	State              uint8    `field:"state"`              // 状态
	Http2Enabled       bool     `field:"http2Enabled"`       // 是否支持HTTP/2
}

type OriginOperator struct {
	Id                 any // ID
	AdminId            any // 管理员ID
	UserId             any // 用户ID
	ReverseProxyId     any // 所属反向代理ID
	IsOn               any // 是否启用
	Name               any // 名称
	Version            any // 版本
	Addr               any // 地址
	Oss                any // OSS配置
	Description        any // 描述
	Code               any // 代号
	Weight             any // 权重
	ConnTimeout        any // 连接超时
	ReadTimeout        any // 读超时
	IdleTimeout        any // 空闲连接超时
	MaxFails           any // 最多失败次数
	MaxConns           any // 最大并发连接数
	MaxIdleConns       any // 最多空闲连接数
	HttpRequestURI     any // 转发后的请求URI
	HttpRequestHeader  any // 请求Header配置
	HttpResponseHeader any // 响应Header配置
	Host               any // 自定义主机名
	HealthCheck        any // 健康检查设置
	Cert               any // 证书设置
	Ftp                any // FTP相关设置
	CreatedAt          any // 创建时间
	Domains            any // 所属域名
	FollowPort         any // 端口跟随
	State              any // 状态
	Http2Enabled       any // 是否支持HTTP/2
}

func NewOriginOperator() *OriginOperator {
	return &OriginOperator{}
}
