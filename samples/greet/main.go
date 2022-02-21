package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func post(id int) {
	for {
		fmt.Println("Hello world ID#", id)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go post(102)
	go post(101)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	<-stopChan

	fmt.Println("Graceful shutdown")
}
