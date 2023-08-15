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
		tokens, _ := scanning.Scan(sourceCode)

		tokens = scanning.ProcessLeadingTrailingDocumentation(sourceCode, tokens)

		expression := ParseExpression(sourceCode, tokens)

		assert.Equal(t, sExpression, SExpression(expression), "For source code: "+sourceCode)
	}

	t.Run("identifier literals", func(t *testing.T) {
		check("abc", "(id abc)")
		check("\n  d  \n", "(id d)")
	})

	t.Run("integer literals", func(t *testing.T) {
		check("123", "(int 123)")
		check("789", "(int 789)")
	})

	t.Run("floating point literals", func(t *testing.T) {
		check("1.23", "(float 1.23)")
		check("78.9", "(float 78.9)")
	})

	t.Run("multiline string literals", func(t *testing.T) {
		check("` line one\n ` line two\n", "(multilinestr\n` line one\n ` line two\n)")
	})

	t.Run("string literals", func(t *testing.T) {
		check(`"123"`, `(string "123")`)
		check(`'789'`, `(string '789')`)
	})

	t.Run("leading documentation", func(t *testing.T) {
		check("// line one\n // line two\nq", "(doc (leadingdoc\n// line one\n // line two\n) (id q))")
	})

	t.Run("trailing documentation", func(t *testing.T) {
		check("q // line one\n // line two\n", "(doc (id q) (trailingdoc\n// line one\n // line two\n))")
	})

	t.Run("addition", func(t *testing.T) {
		check("x + 1", `(add (id x) (int 1))`)
		check(" 3 + y", `(add (int 3) (id y))`)
		check("x + 1.7", `(add (id x) (float 1.7))`)
		check(" 3.666 + y", `(add (float 3.666) (id y))`)
	})

	t.Run("table of expressions", func(t *testing.T) {
		type parseOutcome struct {
			sourceCode  string
			sExpression string
		}

		tests := []parseOutcome{
			{"x + 1", "(add (id x) (int 1))"},
			{"q - 4", "(subtract (id q) (int 4))"},
			{"a - b + 3", "(add (subtract (id a) (id b)) (int 3))"},
			{"a + b + 3", "(add (add (id a) (id b)) (int 3))"},
			{"1 * 2", "(multiply (int 1) (int 2))"},
			{"x + 3 * g", "(add (id x) (multiply (int 3) (id g)))"},
			{"a + b / 2 - c", "(subtract (add (id a) (divide (id b) (int 2))) (id c))"},
			{"-a", "(negate (id a))"},
			{"-2 * a - b * -r", "(subtract (multiply (negate (int 2)) (id a)) (multiply (id b) (negate (id r))))"},
			{"a.b.c", "(fieldref (fieldref (id a) (id b)) (id c))"},
			{"x.y + z.q", "(add (fieldref (id x) (id y)) (fieldref (id z) (id q)))"},
			{"\"s\"", "(string \"s\")"},
			{"\"string tied in a knot\"", "(string \"string tied in a knot\")"},
			{"'c'", "(string 'c')"},

			{"(x + 5)", "(parenthesized \"()\" (add (id x) (int 5)))"},
			{"((x + 5) / 3)", "(parenthesized \"()\" (divide (parenthesized \"()\" (add (id x) (int 5))) (int 3)))"},

			{"()", "(parenthesized \"()\")"},
			{"(x: int && 5)", "(parenthesized \"()\" (qualify (id x) (intersectlowprecedence (id int) (int 5))))"},
			{"(x: int && 5, y: string && \"s\")", "(parenthesized \"()\" (qualify (id x) (intersectlowprecedence (id int) (int 5))) (qualify (id y) (intersectlowprecedence (id string) (string \"s\"))))"},
			{"(1, 2, 3, 4, 5)", "(parenthesized \"()\" (int 1) (int 2) (int 3) (int 4) (int 5))"},
			{"(1.0, 2.0, 3.0, 4.0, 5.0)", "(parenthesized \"()\" (float 1.0) (float 2.0) (float 3.0) (float 4.0) (float 5.0))"},

			{"{}", "(parenthesized \"{}\")"},
			{"{x: int && 5}", "(parenthesized \"{}\" (qualify (id x) (intersectlowprecedence (id int) (int 5))))"},
			{"{x: int && 5, y: string && \"s\"}", "(parenthesized \"{}\" (qualify (id x) (intersectlowprecedence (id int) (int 5))) (qualify (id y) (intersectlowprecedence (id string) (string \"s\"))))"},
			{"{x: int ?: 5, y: string ?: \"s\"}", "(parenthesized \"{}\" (intersectdefault (qualify (id x) (id int)) (int 5)) (intersectdefault (qualify (id y) (id string)) (string \"s\")))"},
			{"{1, 2, 3, 4, 5}", "(parenthesized \"{}\" (int 1) (int 2) (int 3) (int 4) (int 5))"},

			{"[]", "(sequence)"},
			{"[1, 2, 3, 4, 5]", "(sequence (int 1) (int 2) (int 3) (int 4) (int 5))"},

			{"true and false", "(and (bool true) (bool false))"},
			{"a and b", "(and (id a) (id b))"},
			{"a and b or c", "(or (and (id a) (id b)) (id c))"},
			{"a and not b", "(and (id a) (not (id b)))"},
			{"not a or b", "(or (not (id a)) (id b))"},

			{"1 == 2", "(equals (int 1) (int 2))"},
			{"1 + 1 == 2 / 1", "(equals (add (int 1) (int 1)) (divide (int 2) (int 1)))"},
			{"1 + 1 < 2 / 1", "(less (add (int 1) (int 1)) (divide (int 2) (int 1)))"},
			{"1 + 1 <= 2 / 1", "(lessorequals (add (int 1) (int 1)) (divide (int 2) (int 1)))"},
			{"1 + 1 >= 2 / 1", "(greaterorequals (add (int 1) (int 1)) (divide (int 2) (int 1)))"},

			{"x =~ y", "(match (id x) (id y))"},
			{"x !~ y", "(notmatch (id x) (id y))"},

			{"int?", "(optional (id int))"},
			{"float | int?", "(union (id float) (optional (id int)))"},
			{"float & 7.0", "(intersect (id float) (float 7.0))"},

			{"f(x: 0)", "(call (id f) (parenthesized \"()\" (qualify (id x) (int 0))))"},
			{"(a: f(x: 0))", "(parenthesized \"()\" (qualify (id a) (call (id f) (parenthesized \"()\" (qualify (id x) (int 0))))))"},

			{"1..9", "(range (int 1) (int 9))"},
			{"x in 1..9", "(in (id x) (range (int 1) (int 9)))"},

			{"x is Widget", "(is (id x) (id Widget))"},

			{"1 when n == 0\n| n * f(n-1) when n > 0", "(union (when (int 1) (equals (id n) (int 0))) (when (multiply (id n) (call (id f) (parenthesized \"()\" (subtract (id n) (int 1))))) (greater (id n) (int 0))))"},
			{"f: (n: int) -> int = 1 when n == 0\n| n * f(n-1) when n > 0", "(intersectassign (qualify (id f) (arrow (parenthesized \"()\" (qualify (id n) (id int))) (id int))) (union (when (int 1) (equals (id n) (int 0))) (when (multiply (id n) (call (id f) (parenthesized \"()\" (subtract (id n) (int 1))))) (greater (id n) (int 0)))))"},

			{"x = y + z where {y: 3, z: 5}", "(intersectassign (id x) (where (add (id y) (id z)) (parenthesized \"{}\" (qualify (id y) (int 3)) (qualify (id z) (int 5)))))"},
		}

		for _, test := range tests {
			check(test.sourceCode, test.sExpression)
		}
	})

}

//---------------------------------------------------------------------------------------------------------------------
