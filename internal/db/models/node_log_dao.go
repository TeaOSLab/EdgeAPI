package models

import (
	"github.com/TeaOSLab/EdgeAPI/internal/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/dbs"
	timeutil "github.com/iwind/TeaGo/utils/time"
	"strconv"
	"time"
)

type NodeLogDAO dbs.DAO

const ()

func NewNodeLogDAO() *NodeLogDAO {
	return dbs.NewDAO(&NodeLogDAO{
		DAOObject: dbs.DAOObject{
			DB:     Tea.Env,
			Table:  "edgeNodeLogs",
			Model:  new(NodeLog),
			PkName: "id",
		},
	}).(*NodeLogDAO)
}

var SharedNodeLogDAO *NodeLogDAO

func init() {
	dbs.OnReady(func() {
		SharedNodeLogDAO = NewNodeLogDAO()
	})
}

// 创建日志
func (this *NodeLogDAO) CreateLog(tx *dbs.Tx, nodeRole NodeRole, nodeId int64, level string, tag string, description string, createdAt int64) error {
	op := NewNodeLogOperator()
	op.Role = nodeRole
	op.NodeId = nodeId
	op.Level = level
	op.Tag = tag
	op.Description = description
	op.CreatedAt = createdAt
	op.Day = timeutil.FormatTime("Ymd", createdAt)
	err := this.Save(tx, op)
	return err
}

// 清除超出一定日期的日志
func (this *NodeLogDAO) DeleteExpiredLogs(tx *dbs.Tx, days int) error {
	if days <= 0 {
		return errors.New("invalid days '" + strconv.Itoa(days) + "'")
	}
	date := time.Now().AddDate(0, 0, -days)
	expireDay := timeutil.Format("Ymd", date)
	_, err := this.Query(tx).
		Where("day<=:day").
		Param("day", expireDay).
		Delete()
	return err
}

// 计算节点数量
func (this *NodeLogDAO) CountNodeLogs(tx *dbs.Tx, role string, nodeId int64) (int64, error) {
	return this.Query(tx).
		Attr("nodeId", nodeId).
		Attr("role", role).
		Count()
}

// 列出单页日志
func (this *NodeLogDAO) ListNodeLogs(tx *dbs.Tx, role string, nodeId int64, offset int64, size int64) (result []*NodeLog, err error) {
	_, err = this.Query(tx).
		Attr("nodeId", nodeId).
		Attr("role", role).
		Offset(offset).
		Limit(size).
		Slice(&result).
		DescPk().
		FindAll()
	return
}
