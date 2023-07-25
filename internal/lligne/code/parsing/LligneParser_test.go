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
		check("abc", "(id abc)")
		check("\n  d  \n", "(id d)")
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
		check("// line one\n // line two\nq", "( (leadingdoc\n// line one\n// line two\n) (id q))")
	})

	t.Run("trailing documentation", func(t *testing.T) {
		check("q // line one\n // line two\n", "( (id q) (trailingdoc\n// line one\n// line two\n))")
	})

	t.Run("addition", func(t *testing.T) {
		check("x + 1", `(+ (id x) (int 1))`)
		check(" 3 + y", `(+ (int 3) (id y))`)
	})

	t.Run("table of expressions", func(t *testing.T) {
		type parseOutcome struct {
			sourceCode  string
			sExpression string
		}

		tests := []parseOutcome{
			{"x + 1", "(+ (id x) (int 1))"},
			{"q - 4", "(- (id q) (int 4))"},
			{"a - b + 3", "(+ (- (id a) (id b)) (int 3))"},
			{"a + b + 3", "(+ (id a) (id b) (int 3))"},
			{"1 * 2", "(* (int 1) (int 2))"},
			{"x + 3 * g", "(+ (id x) (* (int 3) (id g)))"},
			{"a + b / 2 - c", "(- (+ (id a) (/ (id b) (int 2))) (id c))"},
			{"-a", "(prefix - (id a))"},
			{"-2 * a - b * -r", "(- (* (prefix - (int 2)) (id a)) (* (id b) (prefix - (id r))))"},
			{"a.b.c", "(. (id a) (id b) (id c))"},
			{"x.y + z.q", "(+ (. (id x) (id y)) (. (id z) (id q)))"},
			{"\"s\"", "(string \"s\")"},
			{"\"string tied in a knot\"", "(string \"string tied in a knot\")"},
			{"'c'", "(string 'c')"},

			{"(x + 5)", "(parenthesized () (+ (id x) (int 5)))"},
			{"((x + 5) / 3)", "(parenthesized () (/ (parenthesized () (+ (id x) (int 5))) (int 3)))"},

			{"()", "(parenthesized ())"},
			{"(x: int && 5)", "(parenthesized () (: (id x) (&& (id int) (int 5))))"},
			{"(x: int && 5, y: string && \"s\")", "(parenthesized () (: (id x) (&& (id int) (int 5))) (: (id y) (&& (id string) (string \"s\"))))"},
			{"(1, 2, 3, 4, 5)", "(parenthesized () (int 1) (int 2) (int 3) (int 4) (int 5))"},

			{"{}", "(parenthesized {})"},
			{"{x: int && 5}", "(parenthesized {} (: (id x) (&& (id int) (int 5))))"},
			{"{x: int && 5, y: string && \"s\"}", "(parenthesized {} (: (id x) (&& (id int) (int 5))) (: (id y) (&& (id string) (string \"s\"))))"},
			{"{x: int ?: 5, y: string ?: \"s\"}", "(parenthesized {} (?: (: (id x) (id int)) (int 5)) (?: (: (id y) (id string)) (string \"s\")))"},
			{"{1, 2, 3, 4, 5}", "(parenthesized {} (int 1) (int 2) (int 3) (int 4) (int 5))"},

			{"[]", "(sequence)"},
			{"[1, 2, 3, 4, 5]", "(sequence (int 1) (int 2) (int 3) (int 4) (int 5))"},

			{"a and b", "(and (id a) (id b))"},
			{"a and b or c", "(or (and (id a) (id b)) (id c))"},
			{"a and not b", "(and (id a) (prefix not (id b)))"},
			{"not a or b", "(or (prefix not (id a)) (id b))"},

			{"1 == 2", "(== (int 1) (int 2))"},
			{"1 + 1 == 2 / 1", "(== (+ (int 1) (int 1)) (/ (int 2) (int 1)))"},
			{"1 + 1 < 2 / 1", "(< (+ (int 1) (int 1)) (/ (int 2) (int 1)))"},
			{"1 + 1 <= 2 / 1", "(<= (+ (int 1) (int 1)) (/ (int 2) (int 1)))"},
			{"1 + 1 >= 2 / 1", "(>= (+ (int 1) (int 1)) (/ (int 2) (int 1)))"},

			{"x =~ y", "(=~ (id x) (id y))"},
			{"x !~ y", "(!~ (id x) (id y))"},

			{"int?", "(optional (id int))"},
			{"float | int?", "(| (id float) (optional (id int)))"},

			{"f(x: 0)", "(call (id f) (parenthesized () (: (id x) (int 0))))"},
			{"(a: f(x: 0))", "(parenthesized () (: (id a) (call (id f) (parenthesized () (: (id x) (int 0))))))"},

			{"1..9", "(.. (int 1) (int 9))"},
			{"x in 1..9", "(in (id x) (.. (int 1) (int 9)))"},

			{"x is Widget", "(is (id x) (id Widget))"},

			{"1 when n == 0\n| n * (n-1) when n > 0", "(| (when (int 1) (== (id n) (int 0))) (when (* (id n) (parenthesized () (- (id n) (int 1)))) (> (id n) (int 0))))"},
			{"f: (n: int) -> int = 1 when n == 0\n| n * (n-1) when n > 0", "(= (: (id f) (-> (parenthesized () (: (id n) (id int))) (id int))) (| (when (int 1) (== (id n) (int 0))) (when (* (id n) (parenthesized () (- (id n) (int 1)))) (> (id n) (int 0)))))"},

			{"x = y + z where {y: 3, z: 5}", "(= (id x) (where (+ (id y) (id z)) (parenthesized {} (: (id y) (int 3)) (: (id z) (int 5)))))"},
		}

		for _, test := range tests {
			check(test.sourceCode, test.sExpression)
		}
	})

}

//---------------------------------------------------------------------------------------------------------------------
