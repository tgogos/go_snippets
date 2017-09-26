package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05.000000-0700"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05.000000-0700"))

	fmt.Printf("%#v\n", time.Now())
}
