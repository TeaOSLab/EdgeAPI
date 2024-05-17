// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package tasks_test

import (
	"github.com/TeaOSLab/EdgeAPI/internal/tasks"
	"github.com/iwind/TeaGo/dbs"
	"testing"
	"time"
)

func TestSSLCertUpdateOCSPTask_Loop(t *testing.T) {
	dbs.NotifyReady()

	var task = tasks.NewSSLCertUpdateOCSPTask(1 * time.Minute)
	err := task.Loop()
	if err != nil {
		t.Fatal(err)
	}
}
