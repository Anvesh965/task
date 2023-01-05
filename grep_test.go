package main

import (
	"fmt"
	"testing"
)

func areSame(result []string, expected []string) (bool, string) {

	if len(result) != len(expected) {
		return false, "len(result) != len(expected)"
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			return false, "Error At : Expected -> " + expected[i] + ",but got -> " + result[i]
		}
	}
	return true, ""

}

func Test_1(t *testing.T) {
	result := Search("hello", "-n", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt:1:hello world",
		"input1.txt:3:hello",
		"input1.txt:4:HELLhello WORLD",
		"input1.txt:5:line with hello",
		"input2.txt:1:hello again",
		"input2.txt:2:hello in file2",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}

func Test_2(t *testing.T) {
	result := Search("hello", "-l", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt",
		"input2.txt",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_3(t *testing.T) {
	result := Search("hello", "-i", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt:hello world",
		"input1.txt:hello",
		"input1.txt:HELLhello WORLD",
		"input1.txt:line with hello",
		"input2.txt:hello again",
		"input2.txt:hello in file2",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_4(t *testing.T) {
	result := Search("hello", "-v", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt:Go programming language",
		"input2.txt:third line",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_5(t *testing.T) {
	result := Search("hello", "-x", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt:hello",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_6(t *testing.T) {
	result := Search("hello", "-v -n", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt:2:Go programming language",
		"input2.txt:3:third line",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_7(t *testing.T) {
	result := Search("hello", "-v -n -l", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt",
		"input2.txt",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_8(t *testing.T) {
	result := Search("hello", "-x -n", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt:3:hello",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_9(t *testing.T) {
	result := Search("anvesh", "-i -v -n -x", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt:1:hello world",
		"input1.txt:2:Go programming language",
		"input1.txt:3:hello",
		"input1.txt:4:HELLhello WORLD",
		"input1.txt:5:line with hello",
		"input2.txt:1:hello again",
		"input2.txt:2:hello in file2",
		"input2.txt:3:third line",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_10(t *testing.T) {
	result := Search("hello", "-i -v -n -x", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt:2:Go programming language",
		"input2.txt:3:third line",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_11(t *testing.T) {
	result := Search("hello", "-i -v -n -x -l", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt",
		"input2.txt",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_12(t *testing.T) {
	result := Search("go", "-n -i -v", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt:1:hello world",
		"input1.txt:3:hello",
		"input1.txt:4:HELLhello WORLD",
		"input1.txt:5:line with hello",
		"input2.txt:1:hello again",
		"input2.txt:2:hello in file2",
		"input2.txt:3:third line",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_13(t *testing.T) {
	result := Search("", "-n -v", []string{"input1.txt", "input2.txt"})
	expected := []string{}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_14(t *testing.T) {
	result := Search("world", "", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt:hello world",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
func Test_15(t *testing.T) {
	result := Search("world", "-i -l", []string{"input1.txt", "input2.txt"})
	expected := []string{
		"input1.txt",
	}

	isCorrect, error := areSame(result, expected)
	if !isCorrect {
		fmt.Println(error)
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
