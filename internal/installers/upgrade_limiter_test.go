// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package installers_test

import (
	"testing"
	"time"

	"github.com/TeaOSLab/EdgeAPI/internal/installers"
	"github.com/TeaOSLab/EdgeAPI/internal/utils/sizes"
	"github.com/TeaOSLab/EdgeCommon/pkg/nodeconfigs"
)

func TestNewUpgradeLimiter(t *testing.T) {
	var limiter = installers.NewUpgradeLimiter()
	limiter.UpdateNodeBytes(nodeconfigs.NodeRoleNode, 1, 1)
	limiter.UpdateNodeBytes(nodeconfigs.NodeRoleNode, 2, 5*sizes.M)
	t.Log("limiter:", limiter)
	t.Log("canUpgrade:", limiter.CanUpgrade())

	time.Sleep(1 * time.Second)
	t.Log("canUpgrade:", limiter.CanUpgrade())
	t.Log("limiter:", limiter)
	limiter.UpdateNodeBytes(nodeconfigs.NodeRoleNode, 2, 4*sizes.M)
	t.Log("canUpgrade:", limiter.CanUpgrade())

	t.Log("limiter:", limiter)
}
