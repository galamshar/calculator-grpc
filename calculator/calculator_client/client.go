package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/galamshar/calculator-grpc/calculator/calculator_pb"
	"google.golang.org/grpc"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var nums []int32
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("One number for get prime number decomposition\nMore than one number for get average\nEnter the num(s) : ")

	numbers, _ := reader.ReadString('\n')

	regex := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	allString := regex.FindAllString(numbers, -1)

	for _, element := range allString {
		temp, err := strconv.ParseInt(element, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, int32(temp))
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calculator_pb.NewCalculatorServiceClient(conn)

	if len(nums) == 1 {
		PrimeNumberDecomposition(c, nums[0])
	} else {
		ComputeAverage(c, nums)
	}
}

func PrimeNumberDecomposition(c calculator_pb.CalculatorServiceClient, num int32) {
	ctx := context.Background()
	req := &calculator_pb.PrimeNumberDecompositionRequest{
		Number: num,
	}

	response, err := c.PrimeNumberDecomposition(ctx, req)

	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition RPC %v", err)
	}

	log.Println(response.Answer)
}

func ComputeAverage(c calculator_pb.CalculatorServiceClient, nums []int32) {
	ctx := context.Background()
	req := &calculator_pb.ComputeAverageRequest{
		Numbers: nums,
	}

	response, err := c.ComputeAverage(ctx, req)

	if err != nil {
		log.Fatalf("error while calling ComputeAverage RPC %v", err)
	}

	log.Println(response.Answer)
}
