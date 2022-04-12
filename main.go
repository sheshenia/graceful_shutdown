package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func run(ctx context.Context) error {
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-c
		fmt.Println("Program halted!")
		signal.Stop(c)
		cancel()
	}()
	if err := run(ctx); err != nil {
		log.Fatalln(err)
	}
}
