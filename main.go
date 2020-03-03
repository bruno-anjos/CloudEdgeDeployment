package main

import (
	"github.com/bruno-anjos/CloudEdgeDeployment/executor"
	"fmt"
)

func main() {
	fmt.Println(executor.GetPodmanVersion())
}
