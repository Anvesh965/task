package main

import (
	"bufio"
	"fmt"
	"grep/regx"
	"log"
	"os"
	"strings"
)

func main() {

	args := os.Args

	if len(args) == 1 {
		getUserInput()
	} else {
		groupInputData(args[1:])
	}
	// result := Search("hello", "-x -i", []string{"input1.txt", "input2.txt"})
	// fmt.Println(result)

}
func getUserInput() {

	for {
		fmt.Println("type exit to Exit")
		fmt.Print("grep ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}
		splittingInput(input)
	}

}

func splittingInput(input string) {
	commands := strings.Split(input, " ")
	groupInputData(commands)
}

func groupInputData(commands []string) {
	var flags string
	var files []string
	pattern := ""
	for i := 0; i < len(commands); i++ {
		if string(commands[i][0]) == "-" {
			flags += commands[i] + " "
		} else if pattern == "" {
			pattern = commands[i]
		} else {
			files = append(files, commands[i])
		}
	}

	if len(files) == 0 {
		log.Fatal("enter some file names")
	}

	results := Search(pattern, flags, files)
	for _, result := range results {
		fmt.Println(result)
		// fmt.Printf("%q,\n", result)
	}

}

func Search(pattern string, flags string, files []string) []string {

	flagsList := strings.Split(strings.TrimSpace(flags), " ")
	results := regx.PerformOperations(pattern, flagsList, files)
	return results

}
