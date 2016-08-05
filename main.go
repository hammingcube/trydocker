package main

import (
	"fmt"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/network"
	"golang.org/x/net/context"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	options := types.ImageListOptions{All: true}
	images, err := cli.ImageList(context.Background(), options)
	if err != nil {
		panic(err)
	}

	for _, c := range images {
		fmt.Println(c.RepoTags)
	}

	config := &container.Config{
		Cmd:   []string{"ls", "-a"},
		Image: "ubuntu",
	}
	resp, err := cli.ContainerCreate(context.Background(), config, &container.HostConfig{}, &network.NetworkingConfig{}, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ID)
	cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	//cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	// cli.ContainerExecCreate(context.Background(), "gcc", config)
	//     ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.ContainerExecCreateResponse, error)
	// }

}
