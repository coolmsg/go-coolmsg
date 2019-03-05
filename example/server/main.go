package main

import (
	"fmt"
	"io"
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
			BootstrapFunc: func(c io.ReadWriteCloser) coolmsg.Object { return &example.RootObject{} },
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
		fmt.Println("Got interrupt signal, closing down gently...")
		_ = l.Close()
		<-c
		fmt.Println("Got second interrupt signal, closing down brutally.")
		os.Exit(1)
	}()

	log.Printf("listening on %s", listenAddr)

	err = s.Serve(l)
	log.Printf("error accepting connection: %s", err)

	_ = l.Close()

	log.Printf("waiting for connections to end gracefully...")
	s.Wait()
}
