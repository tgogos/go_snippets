package main

import (
	"fmt"
	"log"
	"os/exec"
	// "time"
)

func main() {

	//
	// this waits until ping finishes and then prints out the output
	//
	out, err := exec.Command("bash", "-c", "ping -c 4 8.8.8.8").Output()
	check(err)
	fmt.Printf("command output:\n%s\n", out)

	//
	// this prints lines while ping is being executed
	//
	cmd := exec.Command("bash", "-c", "ping -c 4 8.8.8.8")

	pipe, err := cmd.StdoutPipe()
	check(err)

	err = cmd.Start()
	check(err)

	// i := 0
	for {
		// time.Sleep(200 * time.Millisecond)
		// i++
		// fmt.Println(i)
		b := make([]byte, 256, 256)
		n, err := pipe.Read(b)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		} else {
			fmt.Print(string(b[:n]))
		}
	}

}

func check(err error) {
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}
}
