// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package tasks_test

import (
	"testing"
	"time"

	"github.com/TeaOSLab/EdgeAPI/internal/tasks"
)

func TestNewHealthCheckTask(t *testing.T) {
	var task = tasks.NewHealthCheckTask(1 * time.Minute)
	err := task.Loop()
	if err != nil {
		t.Fatal(err)
	}
}
