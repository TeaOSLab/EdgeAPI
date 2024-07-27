package models_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeAPI/internal/db/models"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/iwind/TeaGo/bootstrap"
	"github.com/iwind/TeaGo/dbs"
)

func TestIPLibraryFileDAO_GenerateIPLibrary(t *testing.T) {
	dbs.NotifyReady()

	var tx *dbs.Tx
	err := models.SharedIPLibraryFileDAO.GenerateIPLibrary(tx, 4)
	if err != nil {
		t.Fatal(err)
	}
}
