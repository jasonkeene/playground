package lex_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/parsers/ini/lex"
)

func TestLex(t *testing.T) {
	testCases := map[string]struct {
		input          string
		expectedTokens []lex.Token
	}{
		"empty": {
			input: "",
			expectedTokens: []lex.Token{
				{
					Type: lex.TokenEOF,
				},
			},
		},
		"extra whitespace": {
			input: `
				[section]
				key = val
			`,
			expectedTokens: []lex.Token{
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  lex.TokenLeftBracket,
					Value: "[",
				},
				{
					Type:  lex.TokenSection,
					Value: "section",
				},
				{
					Type:  lex.TokenRightBracket,
					Value: "]",
				},
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  lex.TokenKey,
					Value: "key",
				},
				{
					Type:  lex.TokenEqualSign,
					Value: "=",
				},
				{
					Type:  lex.TokenValue,
					Value: "val",
				},
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type: lex.TokenEOF,
				},
			},
		},
		"normal": {
			input: `

[sectionA]
keyA1=valA1
keyA2=valA2

[sectionB]
keyB1=valB1
keyB2=valB2

`,
			expectedTokens: []lex.Token{
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  lex.TokenLeftBracket,
					Value: "[",
				},
				{
					Type:  lex.TokenSection,
					Value: "sectionA",
				},
				{
					Type:  lex.TokenRightBracket,
					Value: "]",
				},
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  lex.TokenKey,
					Value: "keyA1",
				},
				{
					Type:  lex.TokenEqualSign,
					Value: "=",
				},
				{
					Type:  lex.TokenValue,
					Value: "valA1",
				},
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  lex.TokenKey,
					Value: "keyA2",
				},
				{
					Type:  lex.TokenEqualSign,
					Value: "=",
				},
				{
					Type:  lex.TokenValue,
					Value: "valA2",
				},
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  lex.TokenLeftBracket,
					Value: "[",
				},
				{
					Type:  lex.TokenSection,
					Value: "sectionB",
				},
				{
					Type:  lex.TokenRightBracket,
					Value: "]",
				},
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  lex.TokenKey,
					Value: "keyB1",
				},
				{
					Type:  lex.TokenEqualSign,
					Value: "=",
				},
				{
					Type:  lex.TokenValue,
					Value: "valB1",
				},
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  lex.TokenKey,
					Value: "keyB2",
				},
				{
					Type:  lex.TokenEqualSign,
					Value: "=",
				},
				{
					Type:  lex.TokenValue,
					Value: "valB2",
				},
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type:  lex.TokenNewLine,
					Value: "\n",
				},
				{
					Type: lex.TokenEOF,
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			l := lex.NewLexer(tc.input)
			l.Run()

			if !cmp.Equal(l.Tokens, tc.expectedTokens) {
				t.Error(cmp.Diff(l.Tokens, tc.expectedTokens))
			}
		})
	}
}
