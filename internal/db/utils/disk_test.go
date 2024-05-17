// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package dbutils_test

import (
	dbutils "github.com/TeaOSLab/EdgeAPI/internal/db/utils"
	_ "github.com/iwind/TeaGo/bootstrap"
	"testing"
)

func TestHasFreeSpace(t *testing.T) {
	t.Log(dbutils.CheckHasFreeSpace())
	t.Log(dbutils.LocalDatabaseDataDir)
}
