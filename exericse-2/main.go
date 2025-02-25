package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("[ConsoleLogger]:", message)
}

type FileLogger struct{
	fileName string
}

func (f *FileLogger) Log(message string) {
	file, _ := os.OpenFile(f.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	file.WriteString(message + "\n")
}

type Processor struct {
	logger Logger
}

func NewProcessor(logger Logger) *Processor {
	return &Processor{logger: logger}
}

func (p *Processor) Process(task string) {
	p.logger.Log("Proccess task " + task)
}

func main() {
	consoleLogger := &ConsoleLogger{}
	processorWithConsole := NewProcessor(consoleLogger)
	processorWithConsole.Process("Task 1")

	fileLogger := &FileLogger{fileName: "log.txt"}
	processorWithFile := NewProcessor(fileLogger)
	processorWithFile.Process("Task 2")
}
