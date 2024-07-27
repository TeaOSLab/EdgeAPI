package tasks_test

import (
	"testing"
	"time"

	"github.com/TeaOSLab/EdgeAPI/internal/tasks"
	"github.com/iwind/TeaGo/dbs"
)

func TestServerAccessLogCleaner_Loop(t *testing.T) {
	dbs.NotifyReady()

	var task = tasks.NewServerAccessLogCleaner(24 * time.Hour)
	err := task.Loop()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
