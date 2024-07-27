// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package utils

import (
	"os/exec"
	"runtime"

	"github.com/TeaOSLab/EdgeAPI/internal/remotelogs"
	executils "github.com/TeaOSLab/EdgeAPI/internal/utils/exec"
	"github.com/iwind/TeaGo/types"
)

func AddPortsToFirewall(ports []int) {
	for _, port := range ports {
		// Linux
		if runtime.GOOS == "linux" {
			// firewalld
			firewallCmd, _ := executils.LookPath("firewall-cmd")
			if len(firewallCmd) > 0 {
				err := exec.Command(firewallCmd, "--add-port="+types.String(port)+"/tcp").Run()
				if err == nil {
					remotelogs.Println("API_NODE", "add port '"+types.String(port)+"' to firewalld")

					_ = exec.Command(firewallCmd, "--add-port="+types.String(port)+"/tcp", "--permanent").Run()
				}
			}
		}
	}
}
