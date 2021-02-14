package main

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestReadInput(t *testing.T) {
	t.Run("Test ReadInput method without wrap", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("commitzen\n"))
		reader := bufio.NewReader(&stdin)
		result := readInput(reader, 0)
		assertEqual(t, result, "commitzen", "FAILED - ReadInput() doesnt return expected results")
	})

	t.Run("Test ReadInput method with wrap", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("commitzen commitzen commitzen\n"))
		reader := bufio.NewReader(&stdin)
		result := readInput(reader, 20)
		assertEqual(t, result, "commitzen commitzen\ncommitzen", "FAILED - ReadInput() doesnt return expected results")
	})
}

func TestWordwrap(t *testing.T) {
	t.Run("With word wrap = 0", func(t *testing.T) {
		result := wordWrap("commitzen commitzen", 0)
		assertEqual(t, result, "commitzen commitzen", "FAILED - ReadInput() doesnt return expected results")
	})

	t.Run("With word wrap > 0", func(t *testing.T) {
		result := wordWrap("commitzen commitzen", 5)
		assertEqual(t, result, "commitzen\ncommitzen", "FAILED - ReadInput() doesnt return expected results")
	})

	t.Run("With word wrap < len(words)", func(t *testing.T) {
		result := wordWrap("commitzen commitzen", 30)
		assertEqual(t, result, "commitzen commitzen", "FAILED - ReadInput() doesnt return expected results")
	})

	t.Run("With word wrap > len(words)", func(t *testing.T) {
		result := wordWrap("commitzen commitzen commitzen commitzen", 20)
		assertEqual(t, result, "commitzen commitzen\ncommitzen commitzen", "FAILED - ReadInput() doesnt return expected results")
	})
}

func TestCreateMessage(t *testing.T) {
	t.Run("Test Create Type - fix", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("1\n"))
		reader := bufio.NewReader(&stdin)
		result := createMessage(defaultConfig.Type, false, "type", *reader)
		assertEqual(t, result, "fix", "FAILED - CreateMessage() doesnt return expected results")
	})

	t.Run("Test Create Type - feat", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("0\n"))
		reader := bufio.NewReader(&stdin)
		result := createMessage(defaultConfig.Type, false, "type", *reader)
		assertEqual(t, result, "feat", "FAILED - CreateMessage() doesnt return expected results")
	})

	t.Run("Test Create Type - personalized", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("2\n"))
		stdin.Write([]byte("test\n"))
		reader := bufio.NewReader(&stdin)
		result := createMessage(defaultConfig.Type, false, "type", *reader)
		assertEqual(t, result, "test", "FAILED - CreateMessage() doesnt return expected results")
	})

	t.Run("Test Create Type - skip", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("3\n"))
		reader := bufio.NewReader(&stdin)
		result := createMessage(defaultConfig.Type, false, "type", *reader)
		assertEqual(t, result, "", "FAILED - CreateMessage() doesnt return expected results")
	})
}
