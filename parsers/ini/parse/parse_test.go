package parse_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/parsers/ini/parse"
)

func TestParse(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected parse.File
	}{
		"empty": {
			expected: parse.File{
				Name: "test.ini",
				Sections: []parse.Section{
					{},
				},
			},
		},
		"extra whitespace": {
			input: `
				[section]
				key = val
			`,
			expected: parse.File{
				Name: "test.ini",
				Sections: []parse.Section{
					{},
					{
						Name: "section",
						KeyValues: []parse.KeyValue{
							{
								Key:   "key",
								Value: "val",
							},
						},
					},
				},
			},
		},
		"normal": {
			input: `

bareKey=bareValue

[sectionA]
keyA1=valA1
keyA2=valA2

[sectionB]
keyB1=valB1
keyB2=valB2

		`,
			expected: parse.File{
				Name: "test.ini",
				Sections: []parse.Section{
					{
						KeyValues: []parse.KeyValue{
							{
								Key:   "bareKey",
								Value: "bareValue",
							},
						},
					},
					{
						Name: "sectionA",
						KeyValues: []parse.KeyValue{
							{
								Key:   "keyA1",
								Value: "valA1",
							},
							{
								Key:   "keyA2",
								Value: "valA2",
							},
						},
					},
					{
						Name: "sectionB",
						KeyValues: []parse.KeyValue{
							{
								Key:   "keyB1",
								Value: "valB1",
							},
							{
								Key:   "keyB2",
								Value: "valB2",
							},
						},
					},
				},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			f, err := parse.Parse("test.ini", tc.input)

			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(f, tc.expected) {
				t.Error(cmp.Diff(f, tc.expected))
			}
		})
	}
}
