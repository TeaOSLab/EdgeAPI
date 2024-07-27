package models

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSysEvent_DecodeEvent(t *testing.T) {
	event := &SysEvent{
		Type: "serverChange",
	}
	eventObj, err := event.DecodeEvent()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(eventObj)
}
