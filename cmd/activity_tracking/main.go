package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("activity_tracking:", err)
		os.Exit(1)
	}
}
