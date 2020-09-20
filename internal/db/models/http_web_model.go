package models

// HTTP Web
type HTTPWeb struct {
	Id                     uint32 `field:"id"`                     // ID
	IsOn                   uint8  `field:"isOn"`                   // 是否启用
	TemplateId             uint32 `field:"templateId"`             // 模版ID
	AdminId                uint32 `field:"adminId"`                // 管理员ID
	UserId                 uint32 `field:"userId"`                 // 用户ID
	State                  uint8  `field:"state"`                  // 状态
	CreatedAt              uint32 `field:"createdAt"`              // 创建时间
	Root                   string `field:"root"`                   // 资源根目录
	Charset                string `field:"charset"`                // 字符集
	Shutdown               string `field:"shutdown"`               // 临时关闭页面配置
	Pages                  string `field:"pages"`                  // 特殊页面
	RedirectToHttps        string `field:"redirectToHttps"`        // 跳转到HTTPS设置
	Indexes                string `field:"indexes"`                // 首页文件列表
	MaxRequestBodySize     string `field:"maxRequestBodySize"`     // 最大允许的请求内容尺寸
	RequestHeaderPolicyId  uint32 `field:"requestHeaderPolicyId"`  // Request Header策略ID
	ResponseHeaderPolicyId uint32 `field:"responseHeaderPolicyId"` // Response Header策略
	AccessLog              string `field:"accessLog"`              // 访问日志配置
	Stat                   string `field:"stat"`                   // 统计配置
	Gzip                   string `field:"gzip"`                   // Gzip配置
	Cache                  string `field:"cache"`                  // 缓存配置
}

type HTTPWebOperator struct {
	Id                     interface{} // ID
	IsOn                   interface{} // 是否启用
	TemplateId             interface{} // 模版ID
	AdminId                interface{} // 管理员ID
	UserId                 interface{} // 用户ID
	State                  interface{} // 状态
	CreatedAt              interface{} // 创建时间
	Root                   interface{} // 资源根目录
	Charset                interface{} // 字符集
	Shutdown               interface{} // 临时关闭页面配置
	Pages                  interface{} // 特殊页面
	RedirectToHttps        interface{} // 跳转到HTTPS设置
	Indexes                interface{} // 首页文件列表
	MaxRequestBodySize     interface{} // 最大允许的请求内容尺寸
	RequestHeaderPolicyId  interface{} // Request Header策略ID
	ResponseHeaderPolicyId interface{} // Response Header策略
	AccessLog              interface{} // 访问日志配置
	Stat                   interface{} // 统计配置
	Gzip                   interface{} // Gzip配置
	Cache                  interface{} // 缓存配置
}

func NewHTTPWebOperator() *HTTPWebOperator {
	return &HTTPWebOperator{}
}
