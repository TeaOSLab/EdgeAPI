package models

// SSL证书
type SSLCert struct {
	Id          uint32 `field:"id"`          // ID
	AdminId     uint32 `field:"adminId"`     // 管理员ID
	UserId      uint32 `field:"userId"`      // 用户ID
	State       uint8  `field:"state"`       // 状态
	CreatedAt   uint64 `field:"createdAt"`   // 创建时间
	UpdatedAt   uint64 `field:"updatedAt"`   // 修改时间
	IsOn        uint8  `field:"isOn"`        // 是否启用
	Name        string `field:"name"`        // 证书名
	Description string `field:"description"` // 描述
	CertData    string `field:"certData"`    // 证书内容
	KeyData     string `field:"keyData"`     // 密钥内容
	ServerName  string `field:"serverName"`  // 证书使用的主机名
	IsCA        uint8  `field:"isCA"`        // 是否为CA证书
	GroupIds    string `field:"groupIds"`    // 证书分组
	TimeBeginAt uint64 `field:"timeBeginAt"` // 开始时间
	TimeEndAt   uint64 `field:"timeEndAt"`   // 结束时间
	DnsNames    string `field:"dnsNames"`    // DNS名称列表
	CommonNames string `field:"commonNames"` // 发行单位列表
}

type SSLCertOperator struct {
	Id          interface{} // ID
	AdminId     interface{} // 管理员ID
	UserId      interface{} // 用户ID
	State       interface{} // 状态
	CreatedAt   interface{} // 创建时间
	UpdatedAt   interface{} // 修改时间
	IsOn        interface{} // 是否启用
	Name        interface{} // 证书名
	Description interface{} // 描述
	CertData    interface{} // 证书内容
	KeyData     interface{} // 密钥内容
	ServerName  interface{} // 证书使用的主机名
	IsCA        interface{} // 是否为CA证书
	GroupIds    interface{} // 证书分组
	TimeBeginAt interface{} // 开始时间
	TimeEndAt   interface{} // 结束时间
	DnsNames    interface{} // DNS名称列表
	CommonNames interface{} // 发行单位列表
}

func NewSSLCertOperator() *SSLCertOperator {
	return &SSLCertOperator{}
}
