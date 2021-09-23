package generator

import (
	"context"
	"crypto/rand"
	"log"
	"math/big"

	pb "github.com/KinitaL/pass-generator/proto"
	//"time"
)

type Generator struct{}

func (*Generator) Generate(context.Context, *pb.GenRequest) (*pb.GenResponse, error) {
	var numbers []string
	for i := 0; i < 10; i++ {
		i, err := rand.Int(rand.Reader, big.NewInt(9))
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, i.String())
	}

	var password string
	for _, number := range numbers {
		password += number
	}

	return &pb.GenResponse{Password: password}, nil
}
