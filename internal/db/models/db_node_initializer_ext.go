// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .
//go:build !plus

package models

import "github.com/iwind/TeaGo/dbs"

var nsAccessLogDAOMapping = map[int64]any{} // dbNodeId => DAO

func initAccessLogDAO(db *dbs.DB, node *DBNode) {
}
