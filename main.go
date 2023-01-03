package main

import (
	"fmt"
	"task/regx"
)

func main() {

	Search("hello", "-x", []string{"input.txt", "input2.txt"})

}

func handleFlags(pattern string, flags string, files []string) []string {
	switch flags {
	case "-n":
		return regx.LinesWithPattern(pattern, files)
	case "-l":
		return regx.FilesWithMatching(pattern, files)
	case "-i":
		return regx.CaseInSensitiveMatching(pattern, files)
	case "-x":
		return regx.FullLineMatching(pattern, files)
	default:
		return []string{}

	}

}

func Search(pattern string, flags string, files []string) {

	// flagsList := strings.Split(flags, " ")

	res := handleFlags(pattern, flags, files)

	fmt.Println(res)

}
