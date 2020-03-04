package API

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

func PodmanPullImage(conn *varlink.Connection, image string) {
	reply, err := PullImage().Call(conn, image)

	ifExistsPrintErrorAndQuit(err)

	fmt.Println(reply.Id)
	fmt.Println(reply.Logs)
}

func GetPodmanVersion(conn *varlink.Connection) string {
	version, goVersion, commitVersion, builtVersion, osArch, remoteApiVersion, err := GetVersion().Call(conn)

	ifExistsPrintErrorAndQuit(err)

	fmt.Printf("version: %s, golang: %s, commit: %s, built: %s, osArch: %s, remoteAPI: %d\n",
		version, goVersion, commitVersion, builtVersion, osArch, remoteApiVersion)
}

func ifExistsPrintError(err error) {
	if err != nil {
		fmt.Errorf("%s", err)
	}
}

func ifExistsPrintErrorAndQuit(err error) {
	if err != nil {
		fmt.Errorf("%s", err)
		os.Exit(1)
	}
}
