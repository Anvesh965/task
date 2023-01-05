package main

import (
	"testing"
)

func areSame(result []string, expected []string) bool {

	if len(result) != len(expected) {
		return false
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expected[i] {
			return false
		}
	}
	return true

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

	// assert.True(t, DeepEqual(result, expected), "The string slices don't match")
	isCorrect := areSame(result, expected)
	if !isCorrect {
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

	// assert.True(t, DeepEqual(result, expected), "The string slices don't match")
	isCorrect := areSame(result, expected)
	if !isCorrect {
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

	// assert.True(t, DeepEqual(result, expected), "The string slices don't match")
	isCorrect := areSame(result, expected)
	if !isCorrect {
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

	// assert.True(t, DeepEqual(result, expected), "The string slices don't match")
	isCorrect := areSame(result, expected)
	if !isCorrect {
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

	// assert.True(t, DeepEqual(result, expected), "The string slices don't match")
	isCorrect := areSame(result, expected)
	if !isCorrect {
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

	// assert.True(t, DeepEqual(result, expected), "The string slices don't match")
	isCorrect := areSame(result, expected)
	if !isCorrect {
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

	// assert.True(t, DeepEqual(result, expected), "The string slices don't match")
	isCorrect := areSame(result, expected)
	if !isCorrect {
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

	// assert.True(t, DeepEqual(result, expected), "The string slices don't match")
	isCorrect := areSame(result, expected)
	if !isCorrect {
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

	// assert.True(t, DeepEqual(result, expected), "The string slices don't match")
	isCorrect := areSame(result, expected)
	if !isCorrect {
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

	// assert.True(t, DeepEqual(result, expected), "The string slices don't match")
	isCorrect := areSame(result, expected)
	if !isCorrect {
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

	// assert.True(t, DeepEqual(result, expected), "The string slices don't match")
	isCorrect := areSame(result, expected)
	if !isCorrect {
		t.Errorf("FAILED, expected -> %v, got -> %v", expected, result)
	} else {
		t.Logf("SUCCEDED, expected -> %v, got -> %v", expected, result)
	}
}
