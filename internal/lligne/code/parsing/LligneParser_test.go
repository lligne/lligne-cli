//
// # Tests of the parser for Lligne code
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package parsing

import (
	"github.com/stretchr/testify/assert"
	"lligne-cli/internal/lligne/code/scanning"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestLligneParser(t *testing.T) {

	check := func(sourceCode string, sExpression string) {
		scanner := scanning.NewLligneBufferedScanner(
			scanning.NewLligneDocumentationHandlingScanner(
				sourceCode,
				scanning.NewLligneScanner(sourceCode),
			),
		)
		parser := NewLligneParser(scanner)
		model := parser.ParseExpression()

		assert.Equal(t, sExpression, model.SExpression())
	}

	t.Run("identifier literals", func(t *testing.T) {
		check("abc", "(identifier abc)")
		check("\n  d  \n", "(identifier d)")
	})

	t.Run("integer literals", func(t *testing.T) {
		check("123", "(int 123)")
		check("789", "(int 789)")
	})

	t.Run("multiline string literals", func(t *testing.T) {
		check("` line one\n ` line two\n", "(multilinestr\n` line one\n` line two\n)")
	})

	t.Run("string literals", func(t *testing.T) {
		check(`"123"`, `(string "123")`)
		check(`'789'`, `(string '789')`)
	})

	t.Run("leading documentation", func(t *testing.T) {
		check("// line one\n // line two\nq", "( (leadingdoc\n// line one\n// line two\n) (identifier q))")
	})

	t.Run("trailing documentation", func(t *testing.T) {
		check("q // line one\n // line two\n", "( (identifier q) (trailingdoc\n// line one\n// line two\n))")
	})

	t.Run("addition", func(t *testing.T) {
		check("x + 1", `(+ (identifier x) (int 1))`)
		check(" 3 + y", `(+ (int 3) (identifier y))`)
	})

	t.Run("table of expressions", func(t *testing.T) {
		type parseOutcome struct {
			sourceCode  string
			sExpression string
		}

		tests := []parseOutcome{
			{"x + 1", "(+ (identifier x) (int 1))"},
			{"q - 4", "(- (identifier q) (int 4))"},
			{"a - b + 3", "(+ (- (identifier a) (identifier b)) (int 3))"},
			{"a + b + 3", "(+ (identifier a) (identifier b) (int 3))"},
			{"1 * 2", "(* (int 1) (int 2))"},
			{"x + 3 * g", "(+ (identifier x) (* (int 3) (identifier g)))"},
			{"a + b / 2 - c", "(- (+ (identifier a) (/ (identifier b) (int 2))) (identifier c))"},
			{"-a", "(prefix - (identifier a))"},
			{"-2 * a - b * -r", "(- (* (prefix - (int 2)) (identifier a)) (* (identifier b) (prefix - (identifier r))))"},
			{"a.b.c", "(. (identifier a) (identifier b) (identifier c))"},
			{"x.y + z.q", "(+ (. (identifier x) (identifier y)) (. (identifier z) (identifier q)))"},
			{"\"s\"", "(string \"s\")"},
			{"\"string tied in a knot\"", "(string \"string tied in a knot\")"},
			{"'c'", "(string 'c')"},

			{"(x + 5)", "(parenthesized (+ (identifier x) (int 5)))"},
			{"((x + 5) / 3)", "(parenthesized (/ (parenthesized (+ (identifier x) (int 5))) (int 3)))"},

			{"()", "(parenthesized)"},
			{"(x: int && 5)", "(parenthesized (: (identifier x) (&& (identifier int) (int 5))))"},
			{"(x: int && 5, y: string && \"s\")", "(parenthesized (: (identifier x) (&& (identifier int) (int 5))) (: (identifier y) (&& (identifier string) (string \"s\"))))"},

			{"a and b", "(and (identifier a) (identifier b))"},
			{"a and b or c", "(or (and (identifier a) (identifier b)) (identifier c))"},
			{"a and not b", "(and (identifier a) (prefix not (identifier b)))"},
			{"not a or b", "(or (prefix not (identifier a)) (identifier b))"},

			{"1 == 2", "(== (int 1) (int 2))"},
			{"1 + 1 == 2 / 1", "(== (+ (int 1) (int 1)) (/ (int 2) (int 1)))"},
			{"1 + 1 < 2 / 1", "(< (+ (int 1) (int 1)) (/ (int 2) (int 1)))"},
			{"1 + 1 <= 2 / 1", "(<= (+ (int 1) (int 1)) (/ (int 2) (int 1)))"},
			{"1 + 1 >= 2 / 1", "(>= (+ (int 1) (int 1)) (/ (int 2) (int 1)))"},

			{"x =~ y", "(=~ (identifier x) (identifier y))"},
			{"x !~ y", "(!~ (identifier x) (identifier y))"},

			//{"f(x: 0)", "(call (identifier f) (parenthesized (: (identifier x) (int 0))))"},
			//{"(a: f(x: 0))", "(parenthesized (: (identifier a) (call (identifier f) (parenthesized (: (identifier x) (int 0))))))"},

			{"1..9", "(.. (int 1) (int 9))"},
			{"x in 1..9", "(in (identifier x) (.. (int 1) (int 9)))"},

			{"x is Widget", "(is (identifier x) (identifier Widget))"},
		}

		for _, test := range tests {
			check(test.sourceCode, test.sExpression)
		}
	})

}

//---------------------------------------------------------------------------------------------------------------------
