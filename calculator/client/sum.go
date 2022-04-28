package main

import (
	"context"
	"log"

	pb "github.com/bkim8815/gRPC_streaming/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 1,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}

	log.Printf("SUm: %d\n", res.Result)
}
