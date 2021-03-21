package main

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

var wordWraps = []struct {
	word  string
	wrap int
	result string
}{
	{"commitzen commitzen", 0, "commitzen commitzen"},
	{"commitzen commitzen", 5, "commitzen\ncommitzen"},
	{"commitzen commitzen", 30, "commitzen commitzen"},
	{"commitzen commitzen commitzen commitzen", 20, "commitzen commitzen\ncommitzen commitzen"},
}


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
	for _, wordWrap1 := range wordWraps {
		t.Run(wordWrap1.word, func(t *testing.T) {
			result := wordWrap(wordWrap1.word, wordWrap1.wrap)
			assertEqual(t, result, wordWrap1.result, "FAILED - Wordwrap() doesnt return expected results")	
		})
	}
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


	t.Run("Test Create Scope - no input - required", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("test scope\n"))
		reader := bufio.NewReader(&stdin)
		result := createMessage(defaultConfig.Scope, false, "scope", *reader)
		assertEqual(t, result, "test scope", "FAILED - CreateMessage() doesnt return expected results")
	})

	t.Run("Test Create description - required", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("test description\n"))
		reader := bufio.NewReader(&stdin)
		result := createMessage(defaultConfig.Description, false, "description", *reader)
		assertEqual(t, result, "test description", "FAILED - CreateMessage() doesnt return expected results")
	})

	t.Run("Test Create body - required", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("test body\n"))
		reader := bufio.NewReader(&stdin)
		result := createMessage(defaultConfig.Body, false, "body", *reader)
		assertEqual(t, result, "test body", "FAILED - CreateMessage() doesnt return expected results")
	})

	t.Run("Test Create body - not required", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("\n"))
		reader := bufio.NewReader(&stdin)
		body := Body{
			Wrap:     15,
			Required: false,
		}
		result := createMessage(body, false, "body", *reader)
		assertEqual(t, result, "", "FAILED - CreateMessage() doesnt return expected results")
	})

	t.Run("Test Create footer - no input - added", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("test footer test footer test footer\n"))
		reader := bufio.NewReader(&stdin)
		footer := Footer{
			Wrap:     15,
			Required: false,
		}
		result := createMessage(footer, false, "footer", *reader)
		assertEqual(t, result, "test footer\ntest footer\ntest footer", "FAILED - CreateMessage() doesnt return expected results")
	})

	t.Run("Test Create footer - skipping", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("\n"))
		reader := bufio.NewReader(&stdin)
		footer := Footer{
			Wrap:     15,
			Required: false,
		}
		result := createMessage(footer, false, "footer", *reader)
		assertEqual(t, result, "", "FAILED - CreateMessage() doesnt return expected results")
	})
}
