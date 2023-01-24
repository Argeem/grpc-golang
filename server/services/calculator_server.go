package services

import (
	context "context"
	"fmt"
	"io"
	"time"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {}

func (calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("Hello %v", req.Name)
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}

func Fib(n uint32) uint32 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}

func (calculatorServer) Fibonacci(r *FibonacciRequest, stream Calculator_FibonacciServer) error {
	for n := uint32(0); n <= r.N; n++ {
		result := Fib(n)
		res := FibonacciResponse{
			Result: result,
		}
		stream.Send(&res)
		time.Sleep(time.Second / 2)
	}
	return nil
}

func (calculatorServer) Average(stream Calculator_AverageServer) error {
	sum := 0.0
	count := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count += 1
	}
	res := AverageResponse{
		Result: sum / count,
	}
	return stream.SendAndClose(&res)
}
