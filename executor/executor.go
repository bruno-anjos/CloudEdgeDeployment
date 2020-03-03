package executor

import (
	"github.com/varlink/go/varlink"
)

func Init(address string) (*varlink.Connection, error) {
	conn, err := varlink.NewConnection(address)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetPodmanVersion(conn *varlink.Connection) string {
	version, _, _, _, _, _, _ := GetVersion().Call(conn)
	return version
}