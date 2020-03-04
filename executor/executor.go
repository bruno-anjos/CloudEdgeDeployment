package executor

import (
	"fmt"
	"github.com/varlink/go/varlink"
	"os"
)

func Init(address string) (*varlink.Connection, error) {
	conn, err := varlink.NewConnection(address)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetPodmanVersion(conn *varlink.Connection) string {
	version, goVersion, commitVersion, builtVersion, osArch, remoteApiVersion, err := GetVersion().Call(conn)

	if err != nil {
		fmt.Errorf("%s", err)
		os.Exit(1)
	}

	return fmt.Sprintf("version: %s, golang: %s, commit: %s, built: %s, osArch: %s, remoteAPI: %d",
		version, goVersion, commitVersion, builtVersion, osArch, remoteApiVersion)
}