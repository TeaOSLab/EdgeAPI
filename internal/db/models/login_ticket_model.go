package models

import "github.com/iwind/TeaGo/dbs"

const (
	LoginTicketField_Id        dbs.FieldName = "id"        // ID
	LoginTicketField_ExpiresAt dbs.FieldName = "expiresAt" // 过期时间
	LoginTicketField_Value     dbs.FieldName = "value"     // 票据值
	LoginTicketField_AdminId   dbs.FieldName = "adminId"   // 管理员ID
	LoginTicketField_UserId    dbs.FieldName = "userId"    // 用户ID
	LoginTicketField_Ip        dbs.FieldName = "ip"        // 用户IP
)

// LoginTicket 登录票据
type LoginTicket struct {
	Id        uint64 `field:"id"`        // ID
	ExpiresAt uint64 `field:"expiresAt"` // 过期时间
	Value     string `field:"value"`     // 票据值
	AdminId   uint32 `field:"adminId"`   // 管理员ID
	UserId    uint32 `field:"userId"`    // 用户ID
	Ip        string `field:"ip"`        // 用户IP
}

type LoginTicketOperator struct {
	Id        any // ID
	ExpiresAt any // 过期时间
	Value     any // 票据值
	AdminId   any // 管理员ID
	UserId    any // 用户ID
	Ip        any // 用户IP
}

func NewLoginTicketOperator() *LoginTicketOperator {
	return &LoginTicketOperator{}
}
