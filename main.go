package main

import (
	"fmt"
	"github.com/bruno-anjos/CloudEdgeDeployment/API"
)

func main() {
	address := "unix:/run/user/1000/podman/io.podman"
	conn, err := API.Init(address)

	if err != nil {
		fmt.Errorf("%s", err)
	}

	fmt.Println(API.GetPodmanVersion(conn))
}
