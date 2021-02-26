package main

import (
	"context"
	"fmt"
	"github.com/galamshar/calculator-grpc/calculator/calculator_pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Enter the num(s) : ")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calculator_pb.NewCalculatorServiceClient(conn)

	PrimeNumberDecomposition(c)
}

func PrimeNumberDecomposition(c calculator_pb.CalculatorServiceClient) {
	ctx := context.Background()
	req := &calculator_pb.PrimeNumberDecompositionRequest{
		Number: 120,
	}

	response, err := c.PrimeNumberDecomposition(ctx, req)

	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition RPC %v", err)
	}

	log.Println(response)
}

func ComputeAverage(c calculator_pb.CalculatorServiceClient) {

}
