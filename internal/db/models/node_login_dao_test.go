package models

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iwind/TeaGo/dbs"
)

func TestNodeLoginDAO_FindFrequentPorts(t *testing.T) {
	dbs.NotifyReady()

	ports, err := SharedNodeLoginDAO.FindFrequentPorts(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ports)
}
