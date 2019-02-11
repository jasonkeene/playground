package flbconfig_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jasonkeene/playground/parsers/flbconfig"
)

func TestLex(t *testing.T) {
	testCases := map[string]struct {
		input          string
		expectedTokens []flbconfig.Token
	}{
		"empty": {
			input: "",
			expectedTokens: []flbconfig.Token{
				{
					Type: flbconfig.TokenEOF,
				},
			},
		},
		"extra whitespace": {
			input: `
				[section]      
				key  val       
			`,
			expectedTokens: []flbconfig.Token{
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  flbconfig.TokenLeftBracket,
					Value: "[",
				},
				{
					Type:  flbconfig.TokenSection,
					Value: "section",
				},
				{
					Type:  flbconfig.TokenRightBracket,
					Value: "]",
				},
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  flbconfig.TokenKey,
					Value: "key",
				},
				{
					Type:  flbconfig.TokenValue,
					Value: "val       ",
				},
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type: flbconfig.TokenEOF,
				},
			},
		},
		"normal": {
			input: `

[sectionA]
keyA1 valA1
keyA2 valA2

[sectionB]
keyB1 valB1
keyB2 valB2

`,
			expectedTokens: []flbconfig.Token{
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  flbconfig.TokenLeftBracket,
					Value: "[",
				},
				{
					Type:  flbconfig.TokenSection,
					Value: "sectionA",
				},
				{
					Type:  flbconfig.TokenRightBracket,
					Value: "]",
				},
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  flbconfig.TokenKey,
					Value: "keyA1",
				},
				{
					Type:  flbconfig.TokenValue,
					Value: "valA1",
				},
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  flbconfig.TokenKey,
					Value: "keyA2",
				},
				{
					Type:  flbconfig.TokenValue,
					Value: "valA2",
				},
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  flbconfig.TokenLeftBracket,
					Value: "[",
				},
				{
					Type:  flbconfig.TokenSection,
					Value: "sectionB",
				},
				{
					Type:  flbconfig.TokenRightBracket,
					Value: "]",
				},
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  flbconfig.TokenKey,
					Value: "keyB1",
				},
				{
					Type:  flbconfig.TokenValue,
					Value: "valB1",
				},
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  flbconfig.TokenKey,
					Value: "keyB2",
				},
				{
					Type:  flbconfig.TokenValue,
					Value: "valB2",
				},
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  flbconfig.TokenNewLine,
					Value: "\n",
				},
				{
					Type: flbconfig.TokenEOF,
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			l := flbconfig.NewLexer(tc.input)
			l.Run()

			if !cmp.Equal(l.Tokens, tc.expectedTokens) {
				t.Error(cmp.Diff(l.Tokens, tc.expectedTokens))
			}
		})
	}
}
