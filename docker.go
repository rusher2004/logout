package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/wailsapp/wails"
)

var cli *client.Client
var ctx context.Context

func init() {
	var err error
	cli, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	ctx = context.Background()
}

type Containers struct {
	containers []types.Container
	runtime    *wails.Runtime
}

func (c *Containers) WailsInit(runtime *wails.Runtime) error {
	c.runtime = runtime
	myLog := c.runtime.Log.New("Containers")
	myLog.Info("I'm here")
	return nil
}

func NewContainers() (*Containers, error) {
	result := &Containers{}
	return result, nil
}

// func (m *MyContainer) streamLogs() {
// 	i, err := cli.ContainerLogs(ctx, m.info.ID, types.ContainerLogsOptions{
// 		ShowStderr: true,
// 		ShowStdout: true,
// 		Timestamps: false,
// 		Follow:     true,
// 		Tail:       "40",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	hdr := make([]byte, 8)
// 	for {
// 		_, err := i.Read(hdr)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		var w io.Writer
// 		switch hdr[0] {
// 		case 1:
// 			w = os.Stdout
// 		default:
// 			w = os.Stderr
// 		}
// 		count := binary.BigEndian.Uint32(hdr[4:])
// 		dat := make([]byte, count)
// 		_, err = i.Read(dat)
// 		fmt.Fprint(w, string(dat))
// 	}
// }

// GetContainers gets all running container detaills
func (c *Containers) GetContainers() ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	c.containers = containers
	// fmt.Printf("%+v\n", c)

	c.runtime.Events.Emit("got-containers")
	return c.containers, nil
}

func (c *Containers) StreamLogs() {
	for _, co := range c.containers {
		go c.StreamContainerLogs(co.ID)
	}
}

func (c *Containers) StreamContainerLogs(i string) {
	l, err := cli.ContainerLogs(ctx, i, types.ContainerLogsOptions{
		ShowStderr: true,
		ShowStdout: true,
		Timestamps: false,
		Follow:     true,
		Tail:       "40",
	})
	if err != nil {
		log.Fatal(err)
	}
	hdr := make([]byte, 8)
	for {
		_, err := l.Read(hdr)
		if err != nil {
			log.Fatal(err)
		}
		var w io.Writer
		switch hdr[0] {
		case 1:
			w = os.Stdout
		default:
			w = os.Stderr
		}
		count := binary.BigEndian.Uint32(hdr[4:])
		dat := make([]byte, count)
		_, err = l.Read(dat)
		fmt.Fprint(w, string(dat))
		datString := fmt.Sprint(dat)
		c.runtime.Events.Emit("log-stream-"+i, map[string]string{"log": datString})
	}
}

// ContainerInspect gets detail of single container
// func ContainerInspect(containerID string) string {
// 	res, err := cli.ContainerInspect(ctx, containerID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	thing, err := json.Marshal(res.Config)

// fmt.Printf("%+v\n", thing)
// return "hi"
// json.Unmarshal(container, info)
// 	return string(thing)
// }

// ContainerLogs provides stream of logs
// func ContainerLogs(containerID string) {
// 	fmt.Println(containerID)
// options := types.ContainerLogsOptions{
// 	ShowStderr: true,
// 	ShowStdout: true,
// 	Follow:     true,
// }
// res, err := cli.ContainerLogs(ctx, containerID, options)
// if err != nil {
// 	panic(err)
// }
// fmt.Println(containerID)
// fmt.Printf("%+v\n", reflect.TypeOf(res))

//read the first 8 bytes to ignore the HEADER part from docker container logs
// p := make([]byte, 8)
// res.Read(p)
// content, _ := ioutil.ReadAll(res)

// fmt.Printf("res:\t%+v\n", res)
// fmt.Printf("content:\t%+v\n", content)
// return content

// runtime := wailsRuntime{}

// 	i, err := cli.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{
// 		ShowStderr: true,
// 		ShowStdout: true,
// 		Timestamps: false,
// 		Follow:     true,
// 		Tail:       "40",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	hdr := make([]byte, 8)
// 	for {
// 		_, err := i.Read(hdr)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		var w io.Writer
// 		switch hdr[0] {
// 		case 1:
// 			w = os.Stdout
// 		default:
// 			w = os.Stderr
// 		}
// 		count := binary.BigEndian.Uint32(hdr[4:])
// 		dat := make([]byte, count)
// 		_, err = i.Read(dat)
// 		fmt.Fprint(w, string(dat))
// 	}
// }
