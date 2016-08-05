package main

import (
	"fmt"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/network"
	"golang.org/x/net/context"
	"io"
	"log"
	"os"
	"time"
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
		Cmd:         []string{"sh", "script.sh"},
		Image:       "gcc",
		WorkingDir:  "/app",
		AttachStdin: true,
		OpenStdin:   true,
		StdinOnce:   true,
	}
	hostConfig := &container.HostConfig{
		Binds: []string{
			"/Users/madhavjha/src/github.com/maddyonline/trydocker/test_cpp:/app",
		},
	}
	resp, err := cli.ContainerCreate(context.Background(), config, hostConfig, &network.NetworkingConfig{}, "myapp")
	if err != nil {
		panic(err)
	}
	containerId := resp.ID

	readLogs := func() {
		reader, err := cli.ContainerLogs(context.Background(), containerId, types.ContainerLogsOptions{
			ShowStdout: true,
			Follow:     true,
		})
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, reader)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}

	err = cli.ContainerStart(context.Background(), containerId, types.ContainerStartOptions{})
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	ch := make(chan struct{})
	go func() {
		_, err = cli.ContainerWait(ctx, containerId)
		if err != nil {
			log.Fatal(err)
			ch <- struct{}{}
			return
		}
		ch <- struct{}{}
	}()

	go readLogs()

	for {
		select {
		case <-ch:
			fmt.Println("Done")
			err := cli.ContainerRemove(context.Background(), containerId, types.ContainerRemoveOptions{Force: true})
			if err != nil {
				log.Fatal("Error removing container: %v", err)
			}
			return
		case <-time.After(2 * time.Second):
			fmt.Println("Still going...")
		}
	}
}
