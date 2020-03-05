package main

import (
	"github.com/bruno-anjos/CloudEdgeDeployment/API"
)

func main() {
	address := "unix:/run/user/1000/podman/io.podman"
	conn := API.PodmanInit(address)

	API.CheckDependencies()
	API.GetPodmanVersion(conn)
	API.PodmanPullImage(conn, "redis")
}
