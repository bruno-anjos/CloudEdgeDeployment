package main

import (
	"./API"
)

func main() {
	address := "unix:/run/user/1000/podman/io.podman"
	API.CheckDependencies()
	API.StartPodmanAPI()

	conn := API.PodmanInit(address)
	API.GetPodmanVersion(conn)
	API.PodmanPullImage(conn, "redis")
}
