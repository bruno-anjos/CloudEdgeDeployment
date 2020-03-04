package main

import (
	"fmt"
	"github.com/bruno-anjos/CloudEdgeDeployment/executor"
)

func main() {
	address := "unix:/run/user/1000/podman/io.podman"
	conn, err := executor.Init(address)

	if err != nil {
		fmt.Errorf("%s", err)
	}

	fmt.Println(executor.GetPodmanVersion(conn))
}
