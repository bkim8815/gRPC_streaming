package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/bkim8815/gRPC_streaming/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGteet eVeryone was incoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error whiel calling Greeting stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "maries"},
		{FirstName: "test"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)

		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error whiel receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)

		}

		close(waitc)
	}()

	<-waitc

}
