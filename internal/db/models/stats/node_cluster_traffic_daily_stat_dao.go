package stats

import (
	"github.com/TeaOSLab/EdgeAPI/internal/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iwind/TeaGo/Tea"
	"github.com/iwind/TeaGo/dbs"
	"github.com/iwind/TeaGo/maps"
)

type NodeClusterTrafficDailyStatDAO dbs.DAO

func NewNodeClusterTrafficDailyStatDAO() *NodeClusterTrafficDailyStatDAO {
	return dbs.NewDAO(&NodeClusterTrafficDailyStatDAO{
		DAOObject: dbs.DAOObject{
			DB:     Tea.Env,
			Table:  "edgeNodeClusterTrafficDailyStats",
			Model:  new(NodeClusterTrafficDailyStat),
			PkName: "id",
		},
	}).(*NodeClusterTrafficDailyStatDAO)
}

var SharedNodeClusterTrafficDailyStatDAO *NodeClusterTrafficDailyStatDAO

func init() {
	dbs.OnReady(func() {
		SharedNodeClusterTrafficDailyStatDAO = NewNodeClusterTrafficDailyStatDAO()
	})
}

// 增加流量
func (this *NodeClusterTrafficDailyStatDAO) IncreaseDailyBytes(tx *dbs.Tx, clusterId int64, day string, bytes int64) error {
	if len(day) != 8 {
		return errors.New("invalid day '" + day + "'")
	}
	err := this.Query(tx).
		Param("bytes", bytes).
		InsertOrUpdateQuickly(maps.Map{
			"clusterId": clusterId,
			"day":       day,
			"bytes":     bytes,
		}, maps.Map{
			"bytes": dbs.SQL("bytes+:bytes"),
		})
	if err != nil {
		return err
	}
	return nil
}
