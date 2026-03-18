package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello from Docker!")

		time.Sleep(time.Second * 1)
	}
}
