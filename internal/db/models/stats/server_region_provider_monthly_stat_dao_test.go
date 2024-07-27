package stats

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)


func TestServerRegionProviderMonthlyStatDAO_Clean(t *testing.T) {
	var dao = NewServerRegionProviderMonthlyStatDAO()
	err := dao.Clean(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
