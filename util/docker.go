package util

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io"
	"os"
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

func (mydocker MyDocker) docker_start(cli *client.Client, ctx context.Context, imageName string) {
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	NewMyError("启动成功")
}
func (mydocker MyDocker) docker_stop(cli *client.Client, ctx context.Context, ID string) {

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
