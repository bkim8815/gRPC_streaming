package main

import (
	"context"
	"log"
	"time"

	pb "github.com/bkim8815/gRPC_streaming/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("dolonggreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "marie"},
		{FirstName: "test"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling long greet %v\n", err)

	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from long greet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
