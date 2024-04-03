package main

import (
	"fmt"
	"github.com/zishang520/socket.io/v2/socket"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	c = make(chan int)
)

func sendMessage() {
	c <- 1
	time.Sleep(time.Second)
}

func main() {
	io := socket.NewServer(nil, nil)
	http.Handle("/socket.io/", io.ServeHandler(nil))
	go func() {
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatal("error starting http server")
		}
	}()

	err := io.On("connection", func(clients ...any) {
		client := clients[0].(*socket.Socket)
		err := client.On("event", func(data ...any) {
			fmt.Println(data)

		})
		if err != nil {
			return
		}
		err = client.On("disconnect", func(...any) {
		})
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}

	exit := make(chan struct{})
	SignalC := make(chan os.Signal)

	signal.Notify(SignalC, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range SignalC {
			switch s {
			case os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				close(exit)
				return
			}
		}
	}()

	<-exit
	io.Close(nil)
	os.Exit(0)
}
