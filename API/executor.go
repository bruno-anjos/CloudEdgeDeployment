package API

import (
	"fmt"
	"github.com/varlink/go/varlink"
	"os"
	"os/exec"
)

const (
	podmanCommandName = "podman"
	podmanVersionFlag = "-v"
)

func PodmanInit(address string) *varlink.Connection {
	conn, err := varlink.NewConnection(address)

	ifExistsPrintErrorAndQuit(err)

	return conn
}

func PodmanPullImage(conn *varlink.Connection, image string) {
	reply, err := PullImage().Call(conn, image)

	ifExistsPrintErrorAndQuit(err)

	fmt.Println(reply.Id)
	fmt.Println(reply.Logs)
}

func GetPodmanVersion(conn *varlink.Connection) {
	version, goVersion, commitVersion, builtVersion, osArch, remoteApiVersion, err := GetVersion().Call(conn)

	ifExistsPrintErrorAndQuit(err)

	fmt.Printf("version: %s, golang: %s, commit: %s, built: %s, osArch: %s, remoteAPI: %d\n",
		version, goVersion, commitVersion, builtVersion, osArch, remoteApiVersion)
}

func CheckDependencies() {
	podmanCmd := exec.Command(podmanCommandName, podmanVersionFlag)
	err := podmanCmd.Run()
	ifExistsPrintErrorAndQuit(err)
}

func printError(errorString string) {
	_, _ = fmt.Fprintln(os.Stderr, errorString)
}

func ifExistsPrintError(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}

func ifExistsPrintErrorAndQuit(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
