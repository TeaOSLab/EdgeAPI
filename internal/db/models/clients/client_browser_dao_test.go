package clients_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeAPI/internal/db/models/clients"
	_ "github.com/go-sql-driver/mysql"
)

func TestClientBrowserDAO_CreateBrowser(t *testing.T) {
	var dao = clients.NewClientBrowserDAO()
	err := dao.CreateBrowserIfNotExists(nil, "Hello")
	if err != nil {
		t.Fatal(err)
	}

	err = dao.CreateBrowserIfNotExists(nil, "Hello")
	if err != nil {
		t.Fatal(err)
	}

	err = dao.CreateBrowserIfNotExists(nil, "Hello")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClientBrowserDAO_Clean(t *testing.T) {
	var dao = clients.NewClientBrowserDAO()
	err := dao.Clean(nil, 30)
	if err != nil {
		t.Fatal(err)
	}
}
