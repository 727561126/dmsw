package main

import (
    "os"

    "github.com/docker/docker/client"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/pkg/stdcopy"

    "golang.org/x/net/context"
)
type DockerAction interface {
	get_docker()
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

func (mydocker MyDocker) get_docker() {


}

func main() {
    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
    if err != nil {
        panic(err)
    }

    _, err = cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
    if err != nil {
        panic(err)
    }

    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: "alpine",
        Cmd:   []string{"echo", "hello world"},
    }, nil, nil, "")
    if err != nil {
        panic(err)
    }

    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }

    statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
    select {
    case err := <-errCh:
        if err != nil {
            panic(err)
        }
    case <-statusCh:
    }

    out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
    if err != nil {
        panic(err)
    }

    stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

