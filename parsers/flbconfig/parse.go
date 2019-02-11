package flbconfig

import "errors"

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
	l := NewLexer(input)
	l.Run()

	var (
		requireNewline    bool
		processingSection bool
		section           Section

		processingKV bool
		key          string
		value        string
	)

	for _, t := range l.Tokens {
		if requireNewline {
			if t.Type != TokenNewLine {
				return File{}, errors.New("newline required")
			}
			requireNewline = false
			continue
		}

		switch t.Type {
		case TokenError:
			return File{}, errors.New(t.Value)
		case TokenEOF:
			if processingKV || processingSection {
				return File{}, errors.New("unexpected EOF")
			}
			f.Sections = append(f.Sections, section)
			return f, nil
		case TokenNewLine:
			if processingKV || processingSection {
				return File{}, errors.New("unexpected newline")
			}
		case TokenLeftBracket:
			if processingKV || processingSection {
				return File{}, errors.New("unexpected left bracket")
			}
			f.Sections = append(f.Sections, section)
			processingSection = true
			section = Section{}
		case TokenSection:
			if !processingSection {
				return File{}, errors.New("unexpected section")
			}
			section.Name = t.Value
		case TokenRightBracket:
			if !processingSection {
				return File{}, errors.New("unexpected right bracket")
			}
			processingSection = false
			requireNewline = true
		case TokenKey:
			if processingKV || processingSection {
				return File{}, errors.New("unexpected key")
			}
			processingKV = true
			key = t.Value
		case TokenValue:
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
