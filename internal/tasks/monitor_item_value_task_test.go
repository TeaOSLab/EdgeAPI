// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package tasks_test

import (
	"testing"
	"time"

	"github.com/TeaOSLab/EdgeAPI/internal/tasks"
	"github.com/iwind/TeaGo/dbs"
)

func TestMonitorItemValueTask_Loop(t *testing.T) {
	dbs.NotifyReady()

	var task = tasks.NewMonitorItemValueTask(1 * time.Minute)
	err := task.Loop()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
