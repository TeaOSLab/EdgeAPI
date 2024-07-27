// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package services

import (
	"context"
	"errors"

	"github.com/TeaOSLab/EdgeAPI/internal/db/models"
	"github.com/TeaOSLab/EdgeCommon/pkg/iputils"
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

// LoginTicketService 登录票据相关服务
type LoginTicketService struct {
	BaseService
}

// CreateLoginTicket 创建票据
func (this *LoginTicketService) CreateLoginTicket(ctx context.Context, req *pb.CreateLoginTicketRequest) (*pb.CreateLoginTicketResponse, error) {
	_, err := this.ValidateAdmin(ctx)
	if err != nil {
		return nil, err
	}

	if req.AdminId <= 0 && req.UserId <= 0 {
		return nil, errors.New("either 'adminId' or 'userId' must be greater than 0")
	}

	if len(req.Ip) > 0 && !iputils.IsValid(req.Ip) {
		return nil, errors.New("invalid ip: '" + req.Ip + "'")
	}

	var tx = this.NullTx()
	value, err := models.SharedLoginTicketDAO.CreateLoginTicket(tx, req.AdminId, req.UserId, req.Ip)
	if err != nil {
		return nil, err
	}
	return &pb.CreateLoginTicketResponse{Value: value}, nil
}

// FindLoginTicketWithValue 查找票据
// 查找成功后，会自动删除票据信息，所以票据信息只能查询一次
func (this *LoginTicketService) FindLoginTicketWithValue(ctx context.Context, req *pb.FindLoginTicketWithValueRequest) (*pb.FindLoginTicketWithValueResponse, error) {
	_, _, err := this.ValidateAdminAndUser(ctx, false)
	if err != nil {
		return nil, err
	}

	var tx = this.NullTx()
	ticket, err := models.SharedLoginTicketDAO.FindLoginTicketWithValue(tx, req.Value)
	if err != nil {
		return nil, err
	}

	if ticket == nil {
		return &pb.FindLoginTicketWithValueResponse{
			LoginTicket: nil,
		}, nil
	}

	return &pb.FindLoginTicketWithValueResponse{
		LoginTicket: &pb.LoginTicket{
			Id:        int64(ticket.Id),
			ExpiresAt: int64(ticket.ExpiresAt),
			Value:     ticket.Value,
			AdminId:   int64(ticket.AdminId),
			UserId:    int64(ticket.UserId),
			Ip:        ticket.Ip,
		},
	}, nil
}
