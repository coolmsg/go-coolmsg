package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/coolmsg/go-coolmsg"
	"github.com/coolmsg/go-coolmsg/example"
)

func main() {
	listenAddr := "127.0.0.1:4444"
	s := coolmsg.NewServer(coolmsg.ServerOptions{
		ConnOptions: coolmsg.ConnServerOptions{
			// This function returns the bootstrap object, which will be at object id 1 for each connection.
			BootstrapFunc: func() coolmsg.Object { return &example.RootObject{} },
		},
	})

	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("error listening: %s", err)
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		fmt.Println("Got interrupt signal, closing down...")
		l.Close()
		signal.Reset()
	}()

	log.Printf("listening on %s", listenAddr)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("error accepting connection: %s", err)
			break
		}
		s.GoHandle(c)
	}

	_ = l.Close()

	// Wait for any in progress requests to end gracefully.
	log.Printf("waiting for connections to end gracefully...")
	s.Wait()
}
