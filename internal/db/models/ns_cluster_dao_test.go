package models

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/iwind/TeaGo/dbs"
)


func TestNSClusterDAO_DisableNodeCluster(t *testing.T) {
	dbs.NotifyReady()

	err := SharedNSClusterDAO.DisableNSCluster(nil, 7)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
