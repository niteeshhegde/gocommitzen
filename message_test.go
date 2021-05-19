package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
)

var wordWraps = []struct {
	word   string
	wrap   int
	result string
}{
	{"commitzen commitzen", 0, "commitzen commitzen"},
	{"commitzen commitzen", 5, "commitzen\ncommitzen"},
	{"commitzen commitzen", 30, "commitzen commitzen"},
	{"commitzen commitzen commitzen commitzen", 20, "commitzen commitzen\ncommitzen commitzen"},
}

var stdInupts = []struct {
	word   string
	wrap   int
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
	t.Fatal(message, a, b)
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
		want   string
	}{
		"Test Create Type - fix": {
			choise: "1\n",
			input:  "",
			want:   "fix",
		},
		"Test Create Type - feat": {
			choise: "0\n",
			input:  "",
			want:   "feat",
		},
		"Test Create Type - personalized": {
			choise: "2\n",
			input:  "test\n",
			want:   "test",
		},
	}

	scopes := map[string]struct {
		input string
		want  string
	}{
		"Test Create Scope - required": {
			input: "test scope\n",
			want:  "test scope",
		},
		"Test Create Scope - empty": {
			input: "\n",
			want:  "",
		},
	}

	descriptions := map[string]struct {
		input string
		want  string
	}{
		"Test Create description - 1": {
			input: "test description\n",
			want:  "test description",
		},
		"Test Create description - 2": {
			input: "describe123\n",
			want:  "describe123",
		},
	}

	bodies := map[string]struct {
		input string
		want  string
	}{
		"Test Create body - required": {
			input: "test body\n",
			want:  "test body",
		},
		"Test Create body - empty": {
			input: "\n",
			want:  "",
		},
	}

	footers := map[string]struct {
		input string
		want  string
	}{
		"Test Create footer - required": {
			input: "test footer test footer test footer\n",
			want:  "test footer\ntest footer\ntest footer",
		},
		"Test Create footer - empty": {
			input: "\n",
			want:  "",
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
				stdin.Write([]byte(type1.input))
			}
			result := createMessage(defaultConfig.Type, false, "type", *reader)
			assertEqual(t, result, type1.want, "FAILED - CreateMessage() doesnt return expected results for type")
		})
	}

	for name, scope := range scopes {
		scope := scope
		t.Run(name, func(t *testing.T) {
			var stdin bytes.Buffer
			reader := bufio.NewReader(&stdin)
			stdin.Write([]byte(scope.input))
			result := createMessage(defaultConfig.Scope, false, "scope", *reader)
			assertEqual(t, result, scope.want, "FAILED - CreateMessage() doesnt return expected results for scope")
		})
	}

	for name, description := range descriptions {
		description := description
		t.Run(name, func(t *testing.T) {
			var stdin bytes.Buffer
			reader := bufio.NewReader(&stdin)
			stdin.Write([]byte(description.input))
			result := createMessage(defaultConfig.Description, false, "description", *reader)
			assertEqual(t, result, description.want, "FAILED - CreateMessage() doesnt return expected results for description")
		})
	}

	for name, body := range bodies {
		body := body
		t.Run(name, func(t *testing.T) {
			var stdin bytes.Buffer
			reader := bufio.NewReader(&stdin)
			stdin.Write([]byte(body.input))
			result := createMessage(defaultConfig.Body, false, "body", *reader)
			assertEqual(t, result, body.want, "FAILED - CreateMessage() doesnt return expected results for body")
		})
	}

	for name, footer := range footers {
		footer := footer
		t.Run(name, func(t *testing.T) {
			var stdin bytes.Buffer
			reader := bufio.NewReader(&stdin)
			stdin.Write([]byte(footer.input))
			footer1 := Footer{
				Wrap:     15,
				Required: false,
			}
			result := createMessage(footer1, false, "footer", *reader)
			assertEqual(t, result, footer.want, "FAILED - CreateMessage() doesnt return expected results for footer")
		})
	}

}
func BenchmarkConfig(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		newConfig("")
	}
}

func BenchmarkConfig2(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		newConfig("commot.json")
	}
}

func BenchmarkPrint(b *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	b.Run("Error", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			printError("Hello World")
		}
	})
	b.Run("Input", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			printInput("Hello World")
		}
	})
}
func BenchmarkWordWrap(b *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	b.Run("wordWrap 10 - small", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			wordWrap("Hello World Hello World ", 10)
		}
	})
	b.Run("wordWrap 5 - small", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			wordWrap("Hello World Hello World ", 5)
		}
	})
	b.Run("wordWrap 5 - large", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			wordWrap("Hello World Hello World Hello World Hello World ", 5)
		}
	})
	b.Run("wordWrap 10 - large", func(b *testing.B) {
		for i := 0; i <= b.N; i++ {
			wordWrap("Hello World Hello World Hello World Hello World ", 10)
		}
	})
}

func BenchmarkReadInput(b *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	for _, stdInupt := range stdInupts {
		b.Run(stdInupt.word, func(b *testing.B) {
			var stdin bytes.Buffer
			stdin.Write([]byte(stdInupt.word))
			reader := bufio.NewReader(&stdin)
			readInput(reader, stdInupt.wrap)
		})
	}

}

// func TestReadInput(t *testing.T) {
// 	for _, stdInupt := range stdInupts {
// 		t.Run(stdInupt.word, func(t *testing.T) {
// 			var stdin bytes.Buffer
// 			stdin.Write([]byte(stdInupt.word))
// 			reader := bufio.NewReader(&stdin)
// 			result := readInput(reader, stdInupt.wrap)
// 			assertEqual(t, result, stdInupt.result, "FAILED - ReadInput() doesnt return expected results")
// 		})
// 	}
// }

// func BenchmarkInput(b *testing.B) {
// 	b.Run("Input")
// 	for i := 0; i <= b.N; i++ {
// 		printInput("Hello World")
// 	}
// }
