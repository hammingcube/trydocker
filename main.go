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
		Cmd:         []string{"./binary.exe"},
		Image:       "gcc",
		WorkingDir:  "/app",
		AttachStdin: true,
		OpenStdin:   true,
		StdinOnce:   true,
	}
	hostConfig := &container.HostConfig{
		Binds: []string{"/Users/madhavjha/src/github.com/maddyonline/trydocker/test_cpp:/app"},
	}
	resp, err := cli.ContainerCreate(context.Background(), config, hostConfig, &network.NetworkingConfig{}, "")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ID)
	cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})

}
