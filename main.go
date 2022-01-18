package main

import (
	"bufio"
	"fmt"
	"os"
)

func init() {

}

func main() {
	args := os.Args[1:]
	//fmt.Println(len(args))
	if len(args) > 1 {
		fmt.Println("usage: ghuo [script]")
		os.Exit(1)
	}
	if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(file string) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	run(string(bytes))
}

func runPrompt() {
	//line := ""
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, _, err := reader.ReadLine()
		//_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}
		if len(line) > 0 {
			run(string(line))
		}
	}

}

func run(code string) {
	fmt.Println("running", code)
}
