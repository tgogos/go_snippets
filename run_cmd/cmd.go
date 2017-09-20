package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	// "time"
)

func main() {

	ping_approach_1()
	ping_approach_2()
	ping_approach_3()

}

//
// this waits until ping finishes and then prints out the output
//
func ping_approach_1() {
	out, err := exec.Command("bash", "-c", "ping -c 4 8.8.8.8").Output()
	check(err)
	fmt.Printf("command output:\n%s\n", out)
}

//
// this prints lines while ping is being executed
//
func ping_approach_2() {
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

//
// this prints lines while ping is being executed
//
func ping_approach_3() {
	// docker build current directory
	cmdName := "ping"
	cmdArgs := []string{"-c 4", "8.8.8.8"}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("ping -c 4 8.8.8.8 | %s\n", scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}
}
