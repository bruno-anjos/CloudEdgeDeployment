package API

import (
	"bytes"
	"fmt"
	"github.com/varlink/go/varlink"
	"log"
	"os/exec"
)

const (
	podmanCommandName = "podman"
	podmanVersionFlag = "-v"

	varlinkCommandName = "varlink"
	varlinkVersionFlag = "-V"

	podmanAPICommand = "podman"
	podmanAPIVarlinkFlag = "varlink"
	podmanAPITimeoutFlag = "-t 0"
)

func StartPodmanAPI() *varlink.Connection {
	out, err := runCommandWithStdout(podmanAPICommand, podmanAPIVarlinkFlag, podmanAPITimeoutFlag)
	ifExistsPrintErrorAndQuit(err)

	fmt.Printf("Start PodmanAPI: %s\n", out)
}

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
	checkDependency(podmanCommandName, podmanVersionFlag)
	checkDependency(varlinkCommandName, varlinkVersionFlag)
}

func checkDependency(commandName string, args ...string) {
	out, err := runCommandWithStdout(commandName, args...)
	ifExistsPrintErrorAndQuit(err)
	fmt.Printf("[SUCCESS] %s seems to be installed on this machine: %s", commandName, out)
}

func runCommandWithStdout(commandName string, args ...string) (string, error){
	cmd := exec.Command(commandName, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return out.String(), err
}

func ifExistsPrintErrorAndQuit(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
