// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package dnsutils

import (
	"testing"

	"github.com/TeaOSLab/EdgeAPI/internal/db/models"
	"github.com/iwind/TeaGo/dbs"
	"github.com/iwind/TeaGo/logs"
)

func TestNodeClusterDAO_CheckClusterDNS(t *testing.T) {
	dbs.NotifyReady()

	var tx *dbs.Tx
	cluster, err := models.SharedNodeClusterDAO.FindEnabledNodeCluster(tx, 34)
	if err != nil {
		t.Fatal(err)
	}
	if cluster == nil {
		t.Log("cluster not found, skip the test")
		return
	}
	issues, err := CheckClusterDNS(tx, cluster, true)
	if err != nil {
		t.Fatal(err)
	}
	logs.PrintAsJSON(issues, t)
}
