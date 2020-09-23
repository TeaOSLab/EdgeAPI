package models

// 服务
type Server struct {
	Id           uint32 `field:"id"`           // ID
	IsOn         uint8  `field:"isOn"`         // 是否启用
	UniqueId     string `field:"uniqueId"`     // 唯一ID
	UserId       uint32 `field:"userId"`       // 用户ID
	AdminId      uint32 `field:"adminId"`      // 管理员ID
	Type         string `field:"type"`         // 服务类型
	Name         string `field:"name"`         // 名称
	Description  string `field:"description"`  // 描述
	ServerNames  string `field:"serverNames"`  // 域名列表
	Http         string `field:"http"`         // HTTP配置
	Https        string `field:"https"`        // HTTPS配置
	Tcp          string `field:"tcp"`          // TCP配置
	Tls          string `field:"tls"`          // TLS配置
	Unix         string `field:"unix"`         // Unix配置
	Udp          string `field:"udp"`          // UDP配置
	WebId        uint32 `field:"webId"`        // WEB配置
	ReverseProxy string `field:"reverseProxy"` // 反向代理配置
	GroupIds     string `field:"groupIds"`     // 分组ID列表
	Config       string `field:"config"`       // 服务配置，自动生成
	ClusterId    uint32 `field:"clusterId"`    // 集群ID
	IncludeNodes string `field:"includeNodes"` // 部署条件
	ExcludeNodes string `field:"excludeNodes"` // 节点排除条件
	Version      uint32 `field:"version"`      // 版本号
	CreatedAt    uint64 `field:"createdAt"`    // 创建时间
	IsUpdating   uint8  `field:"isUpdating"`   // 是否正在更新
	State        uint8  `field:"state"`        // 状态
}

type ServerOperator struct {
	Id           interface{} // ID
	IsOn         interface{} // 是否启用
	UniqueId     interface{} // 唯一ID
	UserId       interface{} // 用户ID
	AdminId      interface{} // 管理员ID
	Type         interface{} // 服务类型
	Name         interface{} // 名称
	Description  interface{} // 描述
	ServerNames  interface{} // 域名列表
	Http         interface{} // HTTP配置
	Https        interface{} // HTTPS配置
	Tcp          interface{} // TCP配置
	Tls          interface{} // TLS配置
	Unix         interface{} // Unix配置
	Udp          interface{} // UDP配置
	WebId        interface{} // WEB配置
	ReverseProxy interface{} // 反向代理配置
	GroupIds     interface{} // 分组ID列表
	Config       interface{} // 服务配置，自动生成
	ClusterId    interface{} // 集群ID
	IncludeNodes interface{} // 部署条件
	ExcludeNodes interface{} // 节点排除条件
	Version      interface{} // 版本号
	CreatedAt    interface{} // 创建时间
	IsUpdating   interface{} // 是否正在更新
	State        interface{} // 状态
}

func NewServerOperator() *ServerOperator {
	return &ServerOperator{}
}
