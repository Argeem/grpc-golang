package services

import (
	"context"
	"fmt"
)

type calculatorService struct {
	calculatorClient CalculatorClient
}

type CalculatorService interface {
	Hello(name string) error
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
	fmt.Printf("Request : %v", req.Name)
	fmt.Printf("Response ; %v", res.Result)
	return nil
}
