package main

import (
	"fmt"
	"log"

	pb "github.com/bkim8815/gRPC_streaming/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManytimes function was invoked: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
