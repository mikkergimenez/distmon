package docker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

type Docker struct {
	Containers []types.Container
}

func (d *Docker) Get() Docker {
  return Docker{
		Containers: d.List(),
  }
}

func (d *Docker) List() []types.Container {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

  return containers
}
