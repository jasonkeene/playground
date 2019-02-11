package flbconfig_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jasonkeene/playground/parsers/flbconfig"
)

func TestParse(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected flbconfig.File
		err      bool
	}{
		"empty": {
			expected: flbconfig.File{
				Name: "test.config",
				Sections: []flbconfig.Section{
					{},
				},
			},
		},
		"extra whitespace": {
			input: `
				[section]      
				key  val       
			`,
			expected: flbconfig.File{
				Name: "test.config",
				Sections: []flbconfig.Section{
					{},
					{
						Name: "section",
						KeyValues: []flbconfig.KeyValue{
							{
								Key:   "key",
								Value: "val       ",
							},
						},
					},
				},
			},
		},
		"normal": {
			input: `

bareKey bareValue

[sectionA]
keyA1 valA1
keyA2 valA2

[sectionB]
keyB1 valB1
keyB2 valB2

		`,
			expected: flbconfig.File{
				Name: "test.config",
				Sections: []flbconfig.Section{
					{
						KeyValues: []flbconfig.KeyValue{
							{
								Key:   "bareKey",
								Value: "bareValue",
							},
						},
					},
					{
						Name: "sectionA",
						KeyValues: []flbconfig.KeyValue{
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
						KeyValues: []flbconfig.KeyValue{
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
		"two sections on one line": {
			input: `
[sectionA][sectionB]
key val
`,
			err: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			f, err := flbconfig.Parse("test.config", tc.input)

			if !cmp.Equal(f, tc.expected) {
				t.Error(cmp.Diff(f, tc.expected))
			}

			if tc.err && err == nil {
				t.Error("expected err, was nil")
			}
			if !tc.err && err != nil {
				t.Errorf("unexpected err: %s", err)
			}
		})
	}
}
