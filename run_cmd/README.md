# Golang: execute command and display output

## 1. waits for the command to finish

The first approach runs `ping -c 4 8.8.8.8`, and the result is being printed after it finishes:

```go
out, err := exec.Command("bash", "-c", "ping -c 4 8.8.8.8").Output()
check(err)
fmt.Printf("command output:\n%s\n", out)
```

## 2. without waiting... (A)

The second approach runs the same command and prints lines while the execution takes place:

```go
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
```

## 3. without waiting... (B)

The following is a slightly modified version of the example found at ["Shelled-out Commands In Golang"](https://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/) by [Nathan LeClaire](https://nathanleclaire.com/):

```go
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

```

## 4. without waiting... (C)

Found at: [Redirect stdout pipe of child process in Go](https://stackoverflow.com/questions/8875038/redirect-stdout-pipe-of-child-process-in-go).

```go
cmdName := "ping"
cmdArgs := []string{"-c 4", "8.8.8.8"}

cmd := exec.Command(cmdName, cmdArgs...)
cmd.Stdout = os.Stdout
cmd.Stderr = os.Stderr
cmd.Run()
```