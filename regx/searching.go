package regx

import (
	"strings"
)


	// if ans {
	// 	if inv {
	// 		temp[3] = true
	// 		fileLinesCount[file]--
	// 	}
	// } else if !inv {
	// 	temp[3] = true
	// 	fileLinesCount[file]--
	// }

func update(lineData interface{}, ans bool, inv bool, file string, fileLinesCount map[string]int) {
	temp := (lineData).([]interface{})
	if ans {
		if inv {
			temp[3] = true
			fileLinesCount[file]--
		}
	} else if !inv {
		temp[3] = true
		fileLinesCount[file]--
	}

}

func LinesWithPattern(pattern string, filesData *[]interface{}, fileLinesCount map[string]int, inv bool) []string {

	output := []string{}

	for _, lineData := range *filesData {
		temp := lineData.([]interface{})
		line, file, isRemoved := temp[0].(string), temp[2].(string), temp[3].(bool)
		if isRemoved {
			continue
		}
		ans := strings.Contains(line, pattern)
		update(lineData, ans, inv, file, fileLinesCount)

	}
	return output

}

func CaseInSensitiveMatching(pattern string, filesData *[]interface{}, fileLinesCount map[string]int, inv bool) []string {
	output := []string{}
	for _, lineData := range *filesData {
		temp := lineData.([]interface{})
		line, file, isRemoved := temp[0].(string), temp[2].(string), temp[3].(bool)
		if isRemoved {
			continue
		}

		line = strings.ToLower(line)

		ans := strings.Contains(line, strings.ToLower(pattern))
		update(lineData, ans, inv, file, fileLinesCount)

	}

	return output
}

func FullLineMatching(pattern string, filesData *[]interface{}, fileLinesCount map[string]int, inv bool) []string {
	output := []string{}

	for _, lineData := range *filesData {
		temp := lineData.([]interface{})
		line, file, isRemoved := temp[0].(string), temp[2].(string), temp[3].(bool)
		if isRemoved {
			continue
		}
		ans := line == pattern
		update(lineData, ans, inv, file, fileLinesCount)

	}

	return output
}
