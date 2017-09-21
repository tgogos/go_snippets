package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

var err error

type Ping struct {
	Count string `json:"count"`
	Ip    string `json:"ip"`
}

func main() {
	http.HandleFunc("/ping", handler)
	http.ListenAndServe(":5050", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		p := Ping{}

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			fmt.Printf("400 Bad request. Problem decoding the received json.\nDetails:\n%s\n", err.Error())
			http.Error(w, err.Error(), 400)
			return
		}

		fmt.Println("POST /ping ", p)
		ping(w, p)

	default:
		http.Error(w, "Only POST is accepted.", 501)
	}
}

func ping(w http.ResponseWriter, a Ping) {

	cmdName := "ping"
	cmdArgs := []string{"-c", a.Count, a.Ip}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		http.Error(w, "Error creating StdoutPipe for Cmd\n"+err.Error(), 500)
		return
	}

	// the following is used to print output of the command
	// as it makes progress...
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
			//
			// TODO:
			// send output to server
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		http.Error(w, "Error starting Cmd\n"+err.Error(), 500)
		return
	}

	// send 200 OK
	fmt.Fprintf(w, "ping started")

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
	}

}
