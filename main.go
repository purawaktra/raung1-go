package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	fmt.Println("Interrupt signal received. Shutting down...")
}
