package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/KinitaL/pass-generator/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":4333", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewPassGeneratorClient(conn)

	fmt.Println("You are connected to server.")
	for {
		fmt.Println("Нажмите, чтобы получить пароль")
		fmt.Scanln()
		res, err := client.Generate(context.Background(), &pb.GenRequest{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.GetPassword())

	}
}
