package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"task/regx"
)

func main() {

	getUserInput()

	// Search("hello", "-n", []string{"input.txt", "input2.txt"})

}

func splittingInput(input string) (string, string, []string) {
	var flags string
	var files []string
	commands := strings.Split(input, " ")
	pattern := commands[0]
	for i := 1; i < len(commands); i++ {
		if string(commands[i][0]) == "-" {
			flags += commands[i]
		} else {
			files = append(files, commands[i])
		}
	}

	return pattern, flags, files
}

func getUserInput() {

	for {
		fmt.Print("grep ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if input == "e" {
			break
		}
		pattern, flags, files := splittingInput(input)

		Search(pattern, flags, files)
		// fmt.Println(result)

	}

}

func handleFlags(pattern string, flags string, files []string) []string {
	switch flags {
	case "-n":
		return regx.LinesWithPattern(pattern, files, false)
	case "-l":
		return regx.FilesWithMatching(pattern, files)
	case "-i":
		return regx.CaseInSensitiveMatching(pattern, files)
	case "-x":
		return regx.FullLineMatching(pattern, files)
	case "-v":
		return regx.LinesWithPattern(pattern, files, true)
	default:
		return []string{}

	}

}

func Search(pattern string, flags string, files []string) []string {

	// flagsList := strings.Split(flags, " ")

	res := handleFlags(pattern, flags, files)

	fmt.Println(res)

	return res

}
