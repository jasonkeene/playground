package flbconfig

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

type Token struct {
	Type  TokenType
	Value string
}

//go:generate stringer -type=TokenType

type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF
	TokenNewLine

	TokenLeftBracket
	TokenRightBracket

	TokenSection
	TokenKey
	TokenValue
)

const (
	RuneTab          = '\t'
	RuneSpace        = ' '
	RuneLeftBracket  = '['
	RuneRightBracket = ']'
	RuneNewLine      = '\n'
)

type StateFunc func(*Lexer) StateFunc

type Lexer struct {
	Input  string
	Tokens []Token
	State  StateFunc

	Start int
	Pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		Input: input,
		State: LexStart,
	}
}

func (l *Lexer) Run() {
	for l.State != nil {
		l.State = l.State(l)
	}
}

func (l *Lexer) Emit(tokenType TokenType) {
	l.Tokens = append(l.Tokens, Token{
		Type:  tokenType,
		Value: l.Input[l.Start:l.Pos],
	})
	l.Start = l.Pos
}

func (l *Lexer) Errorf(format string, args ...interface{}) StateFunc {
	l.Tokens = append(l.Tokens, Token{
		Type:  TokenError,
		Value: fmt.Sprintf(format, args...),
	})
	return nil
}

func (l *Lexer) Next() rune {
	result, width := utf8.DecodeRuneInString(l.Input[l.Pos:])
	l.Pos += width
	return result
}

func (l *Lexer) PeekNext() rune {
	result, _ := utf8.DecodeRuneInString(l.Input[l.Pos:])
	return result
}

func (l *Lexer) EOF() bool {
	return l.Pos >= len(l.Input)
}

func LexStart(l *Lexer) StateFunc {
	if l.EOF() {
		return LexEOF
	}

	switch next := l.PeekNext(); next {
	case RuneTab, RuneSpace:
		return LexGlobalWhiteSpace
	case RuneNewLine:
		return LexNewLine
	case RuneLeftBracket:
		return LexLeftBracket
	default:
		return LexKey
	}
}

func LexEOF(l *Lexer) StateFunc {
	l.Emit(TokenEOF)
	return nil
}

func LexGlobalWhiteSpace(l *Lexer) StateFunc {
	l.Next()
	l.Start = l.Pos
	return LexStart
}

func LexKeyWhiteSpace(l *Lexer) StateFunc {
	l.Next()
	l.Start = l.Pos

	next := l.PeekNext()

	if unicode.IsLetter(next) || unicode.IsNumber(next) {
		return LexValue
	}

	switch next {
	case RuneTab, RuneSpace:
		return LexKeyWhiteSpace
	default:
		return l.Errorf("invalid key/val delimiter")
	}
}

func LexNewLine(l *Lexer) StateFunc {
	l.Pos += len(string(RuneNewLine))
	l.Emit(TokenNewLine)
	return LexStart
}

func LexLeftBracket(l *Lexer) StateFunc {
	l.Pos += len(string(RuneLeftBracket))
	l.Emit(TokenLeftBracket)
	return LexSection
}

func LexSection(l *Lexer) StateFunc {
	for {
		if l.EOF() {
			return l.Errorf("unexpected EOF")
		}

		next := l.PeekNext()
		if !unicode.IsLetter(next) && !unicode.IsNumber(next) {
			switch next {
			case RuneRightBracket:
				l.Emit(TokenSection)
				return LexRightBracket
			default:
				return l.Errorf("missing right bracket")
			}
		}

		l.Next()
	}
}

func LexRightBracket(l *Lexer) StateFunc {
	l.Pos += len(string(RuneRightBracket))
	l.Emit(TokenRightBracket)
	return LexStart
}

func LexKey(l *Lexer) StateFunc {
	for {
		if l.EOF() {
			return l.Errorf("unexpected EOF")
		}

		next := l.PeekNext()
		if !unicode.IsLetter(next) && !unicode.IsNumber(next) {
			switch next {
			case RuneTab, RuneSpace:
				l.Emit(TokenKey)
				return LexKeyWhiteSpace
			default:
				return l.Errorf("invalid key")
			}
		}

		l.Next()
	}
}

func LexValue(l *Lexer) StateFunc {
	for {
		if l.EOF() {
			return l.Errorf("unexpected EOF")
		}

		r := l.PeekNext()
		if r == RuneNewLine {
			l.Emit(TokenValue)
			return LexStart
		}

		l.Next()
	}
}
