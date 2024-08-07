package setup

import (
	"testing"

	_ "github.com/iwind/TeaGo/bootstrap"
)

func TestSetup_Run(t *testing.T) {
	setup := NewSetup(&Config{
		APINodeProtocol: "http",
		APINodeHost:     "127.0.0.1",
		APINodePort:     8003,
	})
	err := setup.Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("OK")
}
