package regx

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PerformOperations(pattern string, flags []string, files []string) []string {

	flagsOrder := []string{"-v", "-i", "-n", "-x", "-l"}
	flagMap := map[string]bool{"-v": false, "-i": false, "-n": false, "-x": false, "-l": false}

	zeroFlags := flags[0] == ""
	for _, flag := range flags {
		flagMap[flag] = true
	}

	filesData, linesCountInFiles := createStructuredData(files)

	// handling when there are zero flags or when there is only one flag that is -v
	if (len(flags) == 1 && flagMap["-v"]) || zeroFlags {
		handleFlags(pattern, "-n", &filesData, linesCountInFiles, flagMap["-v"])
	}

	for _, flag := range flagsOrder {
		if !flagMap[flag] || (flagMap[flag] && flag == "-v") {
			continue
		}
		handleFlags(pattern, flag, &filesData, linesCountInFiles, flagMap["-v"])
	}
	return formatOutput(&filesData, linesCountInFiles, flagMap, flagsOrder, files)

}

func formatOutput(filesData *[]interface{}, linesCountInFiles map[string]int, flagMap map[string]bool, flagsOrder []string, files []string) []string {

	output := []string{}
	if flagMap["-l"] {
		for _, file := range files {
			if linesCountInFiles[file] > 0 {
				output = append(output, file)
			}
		}
		return output
	}

	for _, lineData := range *filesData {
		temp := lineData.([]interface{})
		line, num, file, isRemoved := temp[0].(string), temp[1].(int), temp[2].(string), temp[3].(bool)

		if isRemoved {
			continue
		}
		lineNumber, fileName := "", ""
		if flagMap["-n"] {
			lineNumber = strconv.Itoa(num) + ":"
		}
		if len(linesCountInFiles) > 1 {
			fileName = file + ":"
		}

		currResult := fileName + lineNumber + line
		output = append(output, currResult)
	}

	return output

}

func handleFlags(pattern string, flag string, filesData *[]interface{}, linesCountInFilesCount map[string]int, inv bool) []string {

	switch flag {
	case "-n", "-l":
		return LinesWithPattern(pattern, filesData, linesCountInFilesCount, inv)
	case "-i":
		return CaseInSensitiveMatching(pattern, filesData, linesCountInFilesCount, inv)
	case "-x":
		return FullLineMatching(pattern, filesData, linesCountInFilesCount, inv)
	default:
		return []string{}

	}

}

func createStructuredData(files []string) ([]interface{}, map[string]int) {

	eachLine := []interface{}{}
	linesCountInFilesCount := map[string]int{}

	for _, file := range files {
		fileReader(file, &eachLine, linesCountInFilesCount)
	}

	return eachLine, linesCountInFilesCount

}
func fileReader(file string, eachLine *[]interface{}, linesCountInFilesCount map[string]int) {

	myFile, err := os.Open(file)

	if err != nil {
		fmt.Println("unable to read file")
		panic(err)
	}

	scanner := bufio.NewScanner(myFile)
	scanner.Split(bufio.ScanLines)

	defer myFile.Close()

	lineNumber := 1
	for scanner.Scan() {
		line := []interface{}{scanner.Text(), lineNumber, file, false}
		*eachLine = append(*eachLine, line)
		linesCountInFilesCount[file] += 1
		lineNumber += 1
	}

}
