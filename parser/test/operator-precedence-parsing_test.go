package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestOperatorPrecedenceParsing(t *testing.T) {
	var tests = []TestPrecedence{
		{
			input:    "-a * b",
			expected: "((-a) * b)",
		},
		{
			input:    "!-a",
			expected: "(!(-a))",
		},
		{
			input:    "a + b + c",
			expected: "((a + b) + c)",
		},
		{
			input:    "a + b - c",
			expected: "((a + b) - c)",
		},
		{
			input:    "a * b * c",
			expected: "((a * b) * c)",
		},
		{
			input:    "a * b / c",
			expected: "((a * b) / c)",
		},
		{
			input:    "a + b / c",
			expected: "(a + (b / c))",
		},
		{
			input:    "a + b * c + d / e - f",
			expected: "(((a + (b * c)) + (d / e)) - f)",
		},
		{
			input:    "3 + 4; -5 * 5",
			expected: "(3 + 4)((-5) * 5)",
		},
		{
			input:    "5 > 4 == 3 < 4",
			expected: "((5 > 4) == (3 < 4))",
		},
		{
			input:    "5 > 4 != 3 < 4",
			expected: "((5 > 4) != (3 < 4))",
		},
		{
			input:    "3 + 4 * 5 == 3 * 1 + 4 * 5",
			expected: "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			input:    "true",
			expected: "true",
		},
		{
			input:    "false",
			expected: "false",
		},
		{
			input:    "3 > 5 == false",
			expected: "((3 > 5) == false)",
		},
		{
			input:    "3 < 5 == true",
			expected: "((3 < 5) == true)",
		},
		{
			input:    "1 + (2 + 3) + 4",
			expected: "((1 + (2 + 3)) + 4)",
		},
		{
			input:    "(5 + 5) * 2",
			expected: "((5 + 5) * 2)",
		},
		{
			input:    "2 / (5 + 5)",
			expected: "(2 / (5 + 5))",
		},
		{
			input:    "-(5 + 5)",
			expected: "(-(5 + 5))",
		},
		{
			input:    "!(true == true)",
			expected: "(!(true == true))",
		},
		{
			input:    "a + add(b * c) + d",
			expected: "((a + add((b * c))) + d)",
		},
		{
			input:    "add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			expected: "add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			input:    "add(a + b + c * d / f + g)",
			expected: "add((((a + b) + ((c * d) / f)) + g))",
		},
		{
			input:    "a * [1, 2, 3, 4][b * c] * d",
			expected: "((a * ([1, 2, 3, 4][(b * c)])) * d)",
		},
		{
			input:    "add(a * b[2], b[1], 2 * [1, 2][1])",
			expected: "add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))",
		},
	}

	for _, test := range tests {
		var lexer = lexer.New(test.input)
		var parser = parser.New(lexer)
		var program = parser.ParseProgram()
		checkParserErrors(t, parser)

		var actual = program.String()

		if actual != test.expected {
			t.Errorf(
				"expected=%q, got=%q",
				test.expected,
				actual,
			)
		}
	}
}
