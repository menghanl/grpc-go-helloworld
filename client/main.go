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

	dialTimeout = flag.Duration("dialTimeout", 10*time.Second, "timeout for creating grpc.ClientConn, format: 100ms, or 10s, or 2h")
	rpcTimeout  = flag.Duration("rpcTimeout", time.Second, "timeout for each RPC, format: 100ms, or 10s, or 2h")

	totalTime = flag.Duration("time", time.Hour, "total time the binary runs, format: 100ms, or 10s, or 2h")

	xds = flag.Bool("xds", true, "do xds or not")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	ctx, cancel := context.WithTimeout(context.Background(), *dialTimeout)
	defer cancel()

	addr := *address
	if *xds {
		addr = "xds-experimental:///" + addr
	}

	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	end := time.After(*totalTime)

	for {
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), *rpcTimeout)
		p := new(peer.Peer)
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name}, grpc.Peer(p))
		if err != nil {
			cancel()
			log.Fatalf("could not greet: %v", err)
		}
		cancel()
		log.Printf("Greeting: %s, from %v", r.GetMessage(), p.Addr)

		select {
		case <-end:
			return
		default:
		}
		time.Sleep(time.Second)
	}
}
