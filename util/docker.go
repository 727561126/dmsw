package main

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

type DockerAction interface {
	docker_start()
	docker_stop()
	docker_running_list()
	docker_run_list()
	docker_images()
	docker_images_pull()
	docker_images_push()
	docker_logs()
}

type MyDocker struct {
}

func get_docker() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}
	return cli, err
}

func (mydocker MyDocker) docker_run_list(cli *client.Client, ctx context.Context) ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	return containers, nil

}
func main() {
	ctx := context.Background()
	cli, err := get_docker()
	containers, err := new(MyDocker).docker_run_list(cli, ctx)

	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		fmt.Print(container)
	}
}
