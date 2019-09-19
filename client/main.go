// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/peer"

	_ "google.golang.org/grpc/xds/experimental"
)

var (
	address = flag.String("address", "localhost:50051", "server address, with port")
	name    = flag.String("name", "world", "name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, "xds-experimental:///"+*address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	for {
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		p := new(peer.Peer)
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name}, grpc.Peer(p))
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s, from %v", r.GetMessage(), p.Addr)
		time.Sleep(time.Second)
	}
}
