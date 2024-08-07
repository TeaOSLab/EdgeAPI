package models

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iwind/TeaGo/dbs"
)

func TestOriginServerDAO_ComposeOriginConfig(t *testing.T) {
	var tx *dbs.Tx
	config, err := SharedOriginDAO.ComposeOriginConfig(tx, 1, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config)
}
