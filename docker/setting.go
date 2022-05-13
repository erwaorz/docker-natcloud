package docker

import (
	"context"
	"github.com/docker/docker/client"
)

var cli *client.Client
var clierr error
var ctx context.Context

func Setup() {
	ctx = context.Background()
	cli, clierr = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation()) //cli客户端对象
	if clierr != nil {
		panic(clierr)
	}
}
