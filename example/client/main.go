package main

import (
	"log"
	"net"

	"github.com/coolmsg/go-coolmsg"
	"github.com/coolmsg/go-coolmsg/example"
)

// In this example client we connect, we create a new greeter, say hello, then destroy it.

func main() {
	listenAddr := "127.0.0.1:4444"
	c, err := net.Dial("tcp", listenAddr)
	if err != nil {
		log.Fatalf("error connecting: %s", err)
	}

	client := coolmsg.NewClient(c, coolmsg.ClientOptions{})

	log.Printf("Creating a new greeter named bob by contacting the bootstrap object...")

	reply, err := client.Send(coolmsg.BOOTSTRAP_OBJECT_ID, &example.MakeGreeter{Name: "bob"})
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	greeterId := reply.(*coolmsg.ObjectRef).Id

	log.Printf("Saying hello to our new greeter...")

	reply, err = client.Send(greeterId, &example.Hello{From: "client"})
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	log.Printf("Got a reply from: %s", reply.(*example.Hello).From)

	log.Printf("destroying the greeter...")
	reply, err = client.Send(greeterId, &coolmsg.Clunk{})
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	// Assert we got an Ok message.
	_ = reply.(*coolmsg.Ok)

	log.Printf("closing the connection...")

	client.Close()
}
