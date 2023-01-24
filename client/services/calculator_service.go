package services

import (
	"context"
	"fmt"
	"io"
	"time"
)

type calculatorService struct {
	calculatorClient CalculatorClient
}

type CalculatorService interface {
	Hello(name string) error
	Fibonacci(n uint32) error
	Average(numbers ...float64) error
}

func NewCalculatorService(calculatorClient CalculatorClient) CalculatorService {
	return calculatorService{calculatorClient}
}

func (base calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name: name,
	}
	res, err := base.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Printf("Service : Hello\n")
	fmt.Printf("Request : %v\n", req.Name)
	fmt.Printf("Response : %v\n", res.Result)
	return nil
}

func (base calculatorService) Fibonacci(n uint32) error {
	req := FibonacciRequest{
		N: n,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	stream, err := base.calculatorClient.Fibonacci(ctx, &req)
	if err != nil {
		return nil
	}

	fmt.Printf("Service : Fibonacci\n")
	fmt.Printf("Request : %v\n", req.N)

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Response : %v\n", res.Result)
	}
	return nil
}

func (base calculatorService) Average(numbers ...float64) error {
	stream, err := base.calculatorClient.Average(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Service : Average\n")

	for _, number := range numbers {
		req := AverageRequest{
			Number: number,
		}
		fmt.Printf("Request : %v\n", req.Number)
		stream.Send(&req)
		time.Sleep(time.Second / 2)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	fmt.Printf("Response : %v\n", res.Result)
	return nil
}
