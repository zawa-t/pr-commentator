package golangcilint_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	golangcilint "github.com/zawa-t/pr/src/format/json/golangci-lint"

	"github.com/zawa-t/pr/src/report"
	"github.com/zawa-t/pr/src/test/helper"
)

func Test_golangcilint_MakeContents_Decode(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		type testCase struct {
			name     string
			input    string
			expected *golangcilint.JSON
		}

		tests := []testCase{
			{
				name: "golangcilint.JSONにDecodeできるデータ",
				input: `
{
	"Issues": [
		{
			"FromLinter": "unparam",
			"Text": "test message1",
			"Severity": "",
			"SourceLines": [
				"func connectionNameRules(name *string, fieldName string) *validation.FieldRules {"
			],
			"Replacement": null,
			"Pos": {
				"Filename": "main.go",
				"Offset": 3632,
				"Line": 70,
				"Column": 40
			},
			"ExpectNoLint": false,
			"ExpectedNoLintLinter": ""
		},
		{
			"FromLinter": "unparam",
			"Text": "test message2",
			"Severity": "",
			"SourceLines": [
				"func advertiserIDRules(advertiserID *string, fieldName string) *validation.FieldRules {"
			],
			"Replacement": null,
			"Pos": {
				"Filename": "main.go",
				"Offset": 3978,
				"Line": 71,
				"Column": 40
			},
			"ExpectNoLint": false,
			"ExpectedNoLintLinter": ""
		}
	]
}
`,
				expected: &golangcilint.JSON{
					Issues: []golangcilint.Issue{
						{
							FromLinter:  "unparam",
							Text:        "test message1",
							Severity:    "",
							SourceLines: []string{"func connectionNameRules(name *string, fieldName string) *validation.FieldRules {"},
							Pos: golangcilint.Pos{
								Filename: "main.go",
								Offset:   3632,
								Line:     70,
								Column:   40,
							},
							ExpectNoLint:         false,
							ExpectedNoLintLinter: "",
						},
						{
							FromLinter:  "unparam",
							Text:        "test message2",
							Severity:    "",
							SourceLines: []string{"func advertiserIDRules(advertiserID *string, fieldName string) *validation.FieldRules {"},
							Pos: golangcilint.Pos{
								Filename: "main.go",
								Offset:   3978,
								Line:     71,
								Column:   40,
							},
							ExpectNoLint:         false,
							ExpectedNoLintLinter: "",
						},
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := golangcilint.Decode(strings.NewReader(tt.input))
				if assert.NoError(t, err) {
					assert.Equal(t, tt.expected, got)
				}
			})
		}
	})

	t.Run("異常系", func(t *testing.T) {
		type testCase struct {
			name  string
			input string
		}

		tests := []testCase{
			{
				name:  "Decodeできないデータ",
				input: "xxxxx",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := golangcilint.Decode(strings.NewReader(tt.input))
				assert.Nil(t, got)
				assert.Error(t, err)
			})
		}
	})
}

func Test_golangcilint_MakeContents(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		type input struct {
			alternativeText     *string
			customMessageFormat *string
			issues              []golangcilint.Issue
		}
		type testCase struct {
			name     string
			input    input
			expected []report.Content
		}

		tests := []testCase{
			{
				name: "正常系",
				input: input{
					issues: []golangcilint.Issue{
						{
							FromLinter:  "linter",
							Text:        "something text",
							Severity:    "severity",
							SourceLines: []string{"1", "2", "3"},
							Pos: golangcilint.Pos{
								Filename: "example.go",
								Offset:   1,
								Line:     10,
								Column:   5,
							},
							ExpectNoLint:         true,
							ExpectedNoLintLinter: "expectedNoLintLinter",
						},
					},
				},
				expected: []report.Content{
					{
						ID:        report.NewID("example.go", 10, report.DefaultMessage("example.go", 10, "linter", "something text")),
						Linter:    "linter",
						FilePath:  "example.go",
						LineNum:   10,
						ColumnNum: 5,
						Message:   report.DefaultMessage("example.go", 10, "linter", "something text"),
					},
				},
			},
			{
				name: "正常系",
				input: input{
					alternativeText: helper.ToPtr("alternativeText"),
					issues: []golangcilint.Issue{
						{
							FromLinter:  "linter",
							Text:        "something text",
							Severity:    "severity",
							SourceLines: []string{"1", "2", "3"},
							Pos: golangcilint.Pos{
								Filename: "example.go",
								Offset:   1,
								Line:     10,
								Column:   5,
							},
							ExpectNoLint:         true,
							ExpectedNoLintLinter: "expectedNoLintLinter",
						},
					},
				},
				expected: []report.Content{
					{
						ID:        report.NewID("example.go", 10, report.DefaultMessage("example.go", 10, "linter", "alternativeText")),
						Linter:    "linter",
						FilePath:  "example.go",
						LineNum:   10,
						ColumnNum: 5,
						Message:   report.DefaultMessage("example.go", 10, "linter", "alternativeText"),
					},
				},
			},
			{
				name: "正常系",
				input: input{
					customMessageFormat: helper.ToPtr("customMessageFormat"),
					issues: []golangcilint.Issue{
						{
							FromLinter:  "linter",
							Text:        "something text",
							Severity:    "severity",
							SourceLines: []string{"1", "2", "3"},
							Pos: golangcilint.Pos{
								Filename: "example.go",
								Offset:   1,
								Line:     10,
								Column:   5,
							},
							ExpectNoLint:         true,
							ExpectedNoLintLinter: "expectedNoLintLinter",
						},
					},
				},
				expected: []report.Content{
					{
						ID:        report.NewID("example.go", 10, report.CustomMessage("customMessageFormat")),
						Linter:    "linter",
						FilePath:  "example.go",
						LineNum:   10,
						ColumnNum: 5,
						Message:   report.CustomMessage("customMessageFormat"),
					},
				},
			},
			{
				name: "正常系",
				input: input{
					customMessageFormat: helper.ToPtr("{{.Pos.Filename}}"),
					issues: []golangcilint.Issue{
						{
							FromLinter:  "linter",
							Text:        "something text",
							Severity:    "severity",
							SourceLines: []string{"1", "2", "3"},
							Pos: golangcilint.Pos{
								Filename: "example.go",
								Offset:   1,
								Line:     10,
								Column:   5,
							},
							ExpectNoLint:         true,
							ExpectedNoLintLinter: "expectedNoLintLinter",
						},
					},
				},
				expected: []report.Content{
					{
						ID:        report.NewID("example.go", 10, report.CustomMessage("example.go")),
						Linter:    "linter",
						FilePath:  "example.go",
						LineNum:   10,
						ColumnNum: 5,
						Message:   report.CustomMessage("example.go"),
					},
				},
			},
			{
				name: "alternativeTextとcustomMessageFormatが両方指定された場合",
				input: input{
					alternativeText:     helper.ToPtr("alternativeText"),
					customMessageFormat: helper.ToPtr("{{.Pos.Filename}}"),
					issues: []golangcilint.Issue{
						{
							FromLinter:  "linter",
							Text:        "something text",
							Severity:    "severity",
							SourceLines: []string{"1", "2", "3"},
							Pos: golangcilint.Pos{
								Filename: "example.go",
								Offset:   1,
								Line:     10,
								Column:   5,
							},
							ExpectNoLint:         true,
							ExpectedNoLintLinter: "expectedNoLintLinter",
						},
					},
				},
				expected: []report.Content{
					{
						ID:        report.NewID("example.go", 10, report.CustomMessage("example.go")),
						Linter:    "linter",
						FilePath:  "example.go",
						LineNum:   10,
						ColumnNum: 5,
						Message:   report.CustomMessage("example.go"),
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				contents, err := golangcilint.MakeContents(tt.input.alternativeText, tt.input.customMessageFormat, tt.input.issues)
				if assert.NoError(t, err) {
					assert.Equal(t, tt.expected, contents)
				}
			})
		}
	})

	t.Run("異常系", func(t *testing.T) {
		type input struct {
			alternativeText     *string
			customMessageFormat *string
			issues              []golangcilint.Issue
		}
		type testCase struct {
			name  string
			input input
		}

		tests := []testCase{
			{
				name: "指定されたcustomMessageFormatに評価できないField名が指定された場合",
				input: input{
					customMessageFormat: helper.ToPtr("{{.Something}}"),
					issues: []golangcilint.Issue{
						{
							FromLinter:  "linter",
							Text:        "something text",
							Severity:    "severity",
							SourceLines: []string{"1", "2", "3"},
							Pos: golangcilint.Pos{
								Filename: "example.go",
								Offset:   1,
								Line:     10,
								Column:   5,
							},
							ExpectNoLint:         true,
							ExpectedNoLintLinter: "expectedNoLintLinter",
						},
					},
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				contents, err := golangcilint.MakeContents(tt.input.alternativeText, tt.input.customMessageFormat, tt.input.issues)
				assert.Nil(t, contents)
				assert.Error(t, err)
			})
		}
	})
}
