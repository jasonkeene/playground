package parse

import (
	"errors"

	"github.com/jasonkeene/playground/parsers/ini/lex"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Section struct {
	Name      string     `json:"name"`
	KeyValues []KeyValue `json:"keyValuePairs"`
}

type File struct {
	Name     string    `json:"name"`
	Sections []Section `json:"sections"`
}

func Parse(name, input string) (File, error) {
	f := File{
		Name: name,
	}
	l := lex.NewLexer(input)
	l.Run()

	var (
		processingSection bool
		section           Section

		processingKV bool
		key          string
		value        string
	)

	for _, t := range l.Tokens {
		switch t.Type {
		case lex.TokenError:
			return File{}, errors.New(t.Value)
		case lex.TokenEOF:
			if processingKV || processingSection {
				return File{}, errors.New("unexpected EOF")
			}
			f.Sections = append(f.Sections, section)
			return f, nil
		case lex.TokenNewLine:
			if processingKV || processingSection {
				return File{}, errors.New("unexpected newline")
			}
		case lex.TokenLeftBracket:
			if processingKV || processingSection {
				return File{}, errors.New("unexpected left bracket")
			}
			f.Sections = append(f.Sections, section)
			processingSection = true
			section = Section{}
		case lex.TokenSection:
			if !processingSection {
				return File{}, errors.New("unexpected section")
			}
			section.Name = t.Value
		case lex.TokenRightBracket:
			if !processingSection {
				return File{}, errors.New("unexpected right bracket")
			}
			processingSection = false
		case lex.TokenKey:
			if processingKV || processingSection {
				return File{}, errors.New("unexpected key")
			}
			processingKV = true
			key = t.Value
		case lex.TokenEqualSign:
			if !processingKV {
				return File{}, errors.New("unexpected equal sign")
			}
		case lex.TokenValue:
			if !processingKV {
				return File{}, errors.New("unexpected value")
			}
			value = t.Value
			section.KeyValues = append(section.KeyValues, KeyValue{
				Key:   key,
				Value: value,
			})
			processingKV = false
			key = ""
			value = ""
		}
	}

	return f, nil
}
