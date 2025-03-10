package runner

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func StartClients(w http.ResponseWriter, r *http.Request) {
	fmt.Println("              ")
	fmt.Println("Starting Clients")
	fmt.Println("              ")

	network := "l15-dev"

	startVanguard("v0.5.1-develop", network)
	startOrchestrator("v0.5.4-develop", network)
	startPandora("v0.5.3-develop", network)
}

func StopClients(w http.ResponseWriter, r *http.Request) {
	cmnd := exec.Command("lukso", "stop")

	if err := cmnd.Start(); err != nil {
		log.Fatal(err)
	}
}

func StartBinary(client string, version string, args []string) {

	fmt.Println("STARTING " + client + "@" + version)
	fmt.Println("/home/rryter/.lukso/downloads/" + client + "/" + version + "/" + client)

	cmnd := exec.Command("/home/rryter/.lukso/downloads/"+client+"/"+version+"/"+client, args...)
	stdout, _ := cmnd.StdoutPipe()
	cmnd.Stderr = cmnd.Stdout

	if startError := cmnd.Start(); startError != nil {
		log.Println("ERROR STARTING " + client + "@" + version)
		log.Fatal(startError)
	}

	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}
}
