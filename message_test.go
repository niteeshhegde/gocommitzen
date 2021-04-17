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

var stdInupts = []struct {
	word  string
	wrap int
	result string
}{
	{"commitzen\n", 0, "commitzen"},
	{"commitzen commitzen commitzen\n", 20, "commitzen commitzen\ncommitzen"},
}



func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message,a,b)
}

func TestReadInput(t *testing.T) {
	for _, stdInupt := range stdInupts {
		t.Run(stdInupt.word, func(t *testing.T) {
			var stdin bytes.Buffer
			stdin.Write([]byte(stdInupt.word))
			reader := bufio.NewReader(&stdin)
			result := readInput(reader, stdInupt.wrap)
			assertEqual(t, result, stdInupt.result, "FAILED - ReadInput() doesnt return expected results")	
		})
	}
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
	types := map[string]struct {
		choise string
		input  string
		want  string
	}{
		"Test Create Type - fix": {
			choise: "1\n",
			input:  "",
			want:  "fix",
		},
		"Test Create Type - feat": {
			choise: "0\n",
			input:  "",
			want:  "feat",
		},
		"Test Create Type - personalized": {
			choise: "2\n",
			input:  "test\n",
			want:  "test",
		},
	}
	t.Parallel()
	for name, type1 := range types {
		type1 := type1
		t.Run(name, func(t *testing.T) {
			var stdin bytes.Buffer
			stdin.Write([]byte(type1.choise))
			reader := bufio.NewReader(&stdin)
			if type1.input != "" {
				stdin.Write([]byte("test\n"))
			}
			result := createMessage(defaultConfig.Type, false, "type", *reader)
			assertEqual(t, result, type1.want, "FAILED - CreateMessage() doesnt return expected results")
		})
	}

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
