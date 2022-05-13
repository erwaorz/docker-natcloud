package docker

import (
	"github.com/docker/docker/api/types"
	"os/exec"
)

func ContainerList() []types.Container {
	containers, clierr := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if clierr != nil {
		panic(clierr)
	}
	return containers
}
func iptablesexist(ip, porto, port, typee string) bool {
	cmd := exec.Command("bash", "-c", "cat /etc/sysconfig/iptables | grep "+ip+":"+port+" | grep \"dport "+porto+" -j DNAT\" | grep "+typee)
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
