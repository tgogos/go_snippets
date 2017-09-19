# Golang: execute command and show results, without waiting to finish

The first approach runs `ping -c 4 8.8.8.8`, and the result is being printed after it finishes:

```go
out, err := exec.Command("bash", "-c", "ping -c 4 8.8.8.8").Output()
check(err)
fmt.Printf("command output:\n%s\n", out)
```

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