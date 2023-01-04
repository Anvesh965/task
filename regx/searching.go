package regx

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func fileReader(file string) []string {

	myFile, _ := os.Open(file)

	scanner := bufio.NewScanner(myFile)
	scanner.Split(bufio.ScanLines)

	defer myFile.Close()

	data := []string{}

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data

}

func LinesWithPattern(pattern string, files []string, isV bool) []string {

	output := []string{}

	prefix := ""

	for _, file := range files {
		fileData := fileReader(file)
		if len(files) > 1 {
			prefix = file + ": "
		}
		for lineNumber, line := range fileData {
			ans := strings.Contains(line, pattern)
			if ans && !isV {
				output = append(output, prefix+strconv.Itoa(lineNumber + 1)+":"+line)
			} else if isV && !ans {
				output = append(output, prefix+line)

			}
		}
	}
	return output

}

func FilesWithMatching(pattern string, files []string) []string {

	output := []string{}

	for _, file := range files {
		fileData := fileReader(file)
		for _, line := range fileData {
			if strings.Contains(line, pattern) {
				output = append(output, file)
				break
			}
		}
	}

	return output
}

func CaseInSensitiveMatching(pattern string, files []string) []string {
	output := []string{}

	prefix := ""

	for _, file := range files {
		fileData := fileReader(file)
		if len(files) > 1 {
			prefix = file + ": "
		}
		for _, line := range fileData {
			for _, word := range strings.Split(line, " ") {
				if strings.EqualFold(word, pattern) {
					output = append(output, prefix+line)
					break
				}
			}
		}
	}

	return output
}

func FullLineMatching(pattern string, files []string) []string {
	output := []string{}
	prefix := ""
	for _, file := range files {
		fileData := fileReader(file)
		if len(files) > 1 {
			prefix = file + ": "
		}
		for _, line := range fileData {
			if strings.EqualFold(line, pattern) {
				output = append(output, prefix+line)
			}
		}
	}

	return output
}
