package main

import (
	"log"
	"net"

	"github.com/KinitaL/pass-generator/generator"
	pb "github.com/KinitaL/pass-generator/proto"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	gen := &generator.Generator{}

	pb.RegisterPassGeneratorServer(server, gen)

	listener, err := net.Listen("tcp", ":4333")
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
