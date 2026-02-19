package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for range 10 {
		dur := time.Duration(rand.Intn(1000)) * time.Millisecond
		fmt.Printf("Sleeping for %v\n", dur)
		// Sleep for a random duration between 0-1000ms
		time.Sleep(dur)
	}
	fmt.Println("Done!")
}
