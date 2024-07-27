package models

import (
	"errors"
	"math/rand"
	"time"

	teaconst "github.com/TeaOSLab/EdgeAPI/internal/const"
	"github.com/TeaOSLab/EdgeAPI/internal/goman"
	"github.com/TeaOSLab/EdgeAPI/internal/remotelogs"
	"github.com/TeaOSLab/EdgeCommon/pkg/iputils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/dbs"
	"github.com/iwind/TeaGo/rands"
	"github.com/iwind/TeaGo/types"
	stringutil "github.com/iwind/TeaGo/utils/string"
)

func init() {
	if !teaconst.IsMain {
		return
	}

	// 清理过期的票据
	var ticker = time.NewTicker(time.Duration(rands.Int(36, 48)) * time.Hour)
	goman.New(func() {
		for range ticker.C {
			err := SharedLoginTicketDAO.CleanExpiredTickets(nil)
			if err != nil {
				remotelogs.Error("LoginTicketDAO", "clean expired tickets failed: "+err.Error())
			}
		}
	})
}

type LoginTicketDAO dbs.DAO

func NewLoginTicketDAO() *LoginTicketDAO {
	return dbs.NewDAO(&LoginTicketDAO{
		DAOObject: dbs.DAOObject{
			DB:     Tea.Env,
			Table:  "edgeLoginTickets",
			Model:  new(LoginTicket),
			PkName: "id",
		},
	}).(*LoginTicketDAO)
}

var SharedLoginTicketDAO *LoginTicketDAO

func init() {
	dbs.OnReady(func() {
		SharedLoginTicketDAO = NewLoginTicketDAO()
	})
}

// CreateLoginTicket 创建票据
func (this *LoginTicketDAO) CreateLoginTicket(tx *dbs.Tx, adminId int64, userId int64, ip string) (ticketValue string, err error) {
	if adminId <= 0 && userId <= 0 {
		err = errors.New("either 'adminId' or 'userId' must be greater than 0")
		return
	}

	if len(ip) > 0 && !iputils.IsValid(ip) {
		err = errors.New("invalid ip: '" + ip + "'")
		return
	}

	ticketValue = stringutil.Md5(types.String(adminId) + "@" + types.String(userId) + types.String(time.Now().UnixNano()) + "@" + types.String(rand.Int63()) + "@" + ip)

	var op = NewLoginTicketOperator()
	op.AdminId = adminId
	op.UserId = userId
	op.ExpiresAt = time.Now().Unix() + 600 /* 10 minutes */
	op.Ip = ip
	op.Value = ticketValue
	err = this.Save(tx, op)
	if err != nil {
		return
	}

	return ticketValue, nil
}

// FindLoginTicketWithValue 查找票据
func (this *LoginTicketDAO) FindLoginTicketWithValue(tx *dbs.Tx, value string) (*LoginTicket, error) {
	if len(value) == 0 {
		return nil, nil
	}

	if len(value) != 32 {
		return nil, nil
	}

	one, err := this.Query(tx).
		Attr("value", value).
		Gt("expiresAt", time.Now().Unix()).
		Find()
	if one == nil || err != nil {
		return nil, err
	}

	var ticket = one.(*LoginTicket)

	// delete the ticket
	err = this.Query(tx).
		Pk(ticket.Id).
		DeleteQuickly()
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

// CleanExpiredTickets 清理过期的票据
func (this *LoginTicketDAO) CleanExpiredTickets(tx *dbs.Tx) error {
	return this.Query(tx).
		Lt("expiresAt", time.Now().Unix()).
		DeleteQuickly()
}
