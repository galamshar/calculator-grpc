package main

import (
	"context"
	"github.com/galamshar/calculator-grpc/calculator/calculator_pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	calculator_pb.UnimplementedCalculatorServiceServer
}

func (s *Server) PrimeNumberDecomposition(ctx context.Context, req *calculator_pb.PrimeNumberDecompositionRequest) (*calculator_pb.PrimeNumberDecompositionResponse, error) {
	var pfs []int32
	number := req.GetNumber()
	for number%2 == 0 {
		pfs = append(pfs, 2)
		number = number / 2
	}

	for i := 3; int32(i*i) <= number; i = i + 2 {
		for number%int32(i) == 0 {
			pfs = append(pfs, int32(i))
			number = number / int32(i)
		}
	}

	if number > 2 {
		pfs = append(pfs, number)
	}

	res := &calculator_pb.PrimeNumberDecompositionResponse{
		Answer: pfs,
	}

	return res, nil
}

func (s *Server) ComputeAverage(ctx context.Context, req *calculator_pb.ComputeAverageRequest) (*calculator_pb.ComputeAverageResponse, error) {
	numbers := req.GetNumbers()
	var sum int32
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	avg := float32(sum) / float32(len(numbers))

	res := &calculator_pb.ComputeAverageResponse{
		Answer: avg,
	}

	return res, nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()

	calculator_pb.RegisterCalculatorServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
