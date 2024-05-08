package models

import "github.com/iwind/TeaGo/dbs"

const (
	HTTPFirewallRuleSetField_Id                 dbs.FieldName = "id"                 // ID
	HTTPFirewallRuleSetField_IsOn               dbs.FieldName = "isOn"               // 是否启用
	HTTPFirewallRuleSetField_Code               dbs.FieldName = "code"               // 代号
	HTTPFirewallRuleSetField_Name               dbs.FieldName = "name"               // 名称
	HTTPFirewallRuleSetField_Description        dbs.FieldName = "description"        // 描述
	HTTPFirewallRuleSetField_CreatedAt          dbs.FieldName = "createdAt"          // 创建时间
	HTTPFirewallRuleSetField_Rules              dbs.FieldName = "rules"              // 规则列表
	HTTPFirewallRuleSetField_Connector          dbs.FieldName = "connector"          // 规则之间的关系
	HTTPFirewallRuleSetField_State              dbs.FieldName = "state"              // 状态
	HTTPFirewallRuleSetField_AdminId            dbs.FieldName = "adminId"            // 管理员ID
	HTTPFirewallRuleSetField_UserId             dbs.FieldName = "userId"             // 用户ID
	HTTPFirewallRuleSetField_Action             dbs.FieldName = "action"             // 执行的动作（过期）
	HTTPFirewallRuleSetField_ActionOptions      dbs.FieldName = "actionOptions"      // 动作的选项（过期）
	HTTPFirewallRuleSetField_Actions            dbs.FieldName = "actions"            // 一组动作
	HTTPFirewallRuleSetField_IgnoreLocal        dbs.FieldName = "ignoreLocal"        // 忽略局域网请求
	HTTPFirewallRuleSetField_IgnoreSearchEngine dbs.FieldName = "ignoreSearchEngine" // 忽略搜索引擎
)

// HTTPFirewallRuleSet 防火墙规则集
type HTTPFirewallRuleSet struct {
	Id                 uint32   `field:"id"`                 // ID
	IsOn               bool     `field:"isOn"`               // 是否启用
	Code               string   `field:"code"`               // 代号
	Name               string   `field:"name"`               // 名称
	Description        string   `field:"description"`        // 描述
	CreatedAt          uint64   `field:"createdAt"`          // 创建时间
	Rules              dbs.JSON `field:"rules"`              // 规则列表
	Connector          string   `field:"connector"`          // 规则之间的关系
	State              uint8    `field:"state"`              // 状态
	AdminId            uint32   `field:"adminId"`            // 管理员ID
	UserId             uint32   `field:"userId"`             // 用户ID
	Action             string   `field:"action"`             // 执行的动作（过期）
	ActionOptions      dbs.JSON `field:"actionOptions"`      // 动作的选项（过期）
	Actions            dbs.JSON `field:"actions"`            // 一组动作
	IgnoreLocal        bool     `field:"ignoreLocal"`        // 忽略局域网请求
	IgnoreSearchEngine bool     `field:"ignoreSearchEngine"` // 忽略搜索引擎
}

type HTTPFirewallRuleSetOperator struct {
	Id                 any // ID
	IsOn               any // 是否启用
	Code               any // 代号
	Name               any // 名称
	Description        any // 描述
	CreatedAt          any // 创建时间
	Rules              any // 规则列表
	Connector          any // 规则之间的关系
	State              any // 状态
	AdminId            any // 管理员ID
	UserId             any // 用户ID
	Action             any // 执行的动作（过期）
	ActionOptions      any // 动作的选项（过期）
	Actions            any // 一组动作
	IgnoreLocal        any // 忽略局域网请求
	IgnoreSearchEngine any // 忽略搜索引擎
}

func NewHTTPFirewallRuleSetOperator() *HTTPFirewallRuleSetOperator {
	return &HTTPFirewallRuleSetOperator{}
}
