package main

import "fmt"

type Output interface {
	Print(message string)
}

type ConsoleOutput struct{}

func (c *ConsoleOutput) Print(message string) {
	fmt.Println(message)
}

type GreeterService struct {
	output Output
}

func NewGreeterService(output Output) *GreeterService {
	return &GreeterService{
		output: output,
	}
}

func (gs *GreeterService) Greet(name string) {
	gs.output.Print("Hello, " + name + "!")
}

func main() {
	consoleOutput := &ConsoleOutput{}
	gs := GreeterService{consoleOutput}
	gs.Greet("naveen")
}
