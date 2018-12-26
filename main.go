package main

import (
	"fmt"
	"log"
	"net"

	insultJMK "github.com/dixneuf19/insult-jmk/grpc"
	"github.com/dixneuf19/insult-jmk/insulter"
	"google.golang.org/grpc"
)

func main() {
	insults := insulter.CreateInsultDict()

	// create a listener on TCP port 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("listened on TCP 50051")
	// create a server instance
	s := insultJMK.Server{
		Insulter: &insults,
	}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	insultJMK.RegisterInsulterServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	log.Println("GRPC server started")

}
