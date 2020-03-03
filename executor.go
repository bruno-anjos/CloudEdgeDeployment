package executor

import (
	"fmt"
	"github.com/varlink/go/varlink"
	"github.com/bruno-anjos/CloudEdgeDeployment/executor/iopodman"
)

type VarlinkCall struct{ varlink.Call }

func Init(address string) (*varlink.Connection, error) {
	conn, err := varlink.NewConnection(nil, address)
	if err != nil {
		return _, err
	}

	return conn, nil
}

func GetPodmanVersion(connection *varlink.Connection) {
	pod, err := iopodman.GetPod().Call(conn, "foobar")
}