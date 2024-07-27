// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .
//go:build !plus

package setup

import (
	"github.com/iwind/TeaGo/dbs"
)

// 检查自建DNS全局设置
func (this *SQLExecutor) checkNS(db *dbs.DB) error {
	return nil
}
