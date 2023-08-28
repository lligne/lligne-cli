//
// # Tests of the parser for Lligne code
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package formatting

import (
	"github.com/stretchr/testify/assert"
	"lligne-cli/internal/lligne/code/parsing"
	"lligne-cli/internal/lligne/code/scanning"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestLligneFormatter(t *testing.T) {

	check := func(sourceCode string) {
		tokens, _ := scanning.Scan(sourceCode)

		tokens = scanning.ProcessLeadingTrailingDocumentation(sourceCode, tokens)

		expression := parsing.ParseExpression(sourceCode, tokens)

		assert.Equal(t, sourceCode, FormatExpr(sourceCode, expression))
	}

	t.Run("identifier literals", func(t *testing.T) {
		check("abc")
		check("d")
		check("a_bc")
		check("ab-c")
	})

	t.Run("integer literals", func(t *testing.T) {
		check("123")
		check("789")
	})

	t.Run("floating point literals", func(t *testing.T) {
		check("1.23")
		check("78.9")
	})

	t.Run("multiline string literals", func(t *testing.T) {
		check("` line one\n ` line two\n")
	})

	t.Run("string literals", func(t *testing.T) {
		check(`"123"`)
		check(`'789'`)
	})

	//t.Run("leading documentation", func(t *testing.T) {
	//	check("// line one\n // line two\nq", "(doc (leadingdoc\n// line one\n // line two\n) (id q))")
	//})
	//
	//t.Run("trailing documentation", func(t *testing.T) {
	//	check("q // line one\n // line two\n", "(doc (id q) (trailingdoc\n// line one\n // line two\n))")
	//})

	t.Run("addition", func(t *testing.T) {
		check("x + 1")
		check("3 + y")
		check("x + 1.7")
		check("3.666 + y")
	})

	t.Run("built in types", func(t *testing.T) {
		check("x: Int64")
		check("isWorking: Bool")
		check("amount: Float64")
		check("name: String")
	})

	t.Run("table of expressions", func(t *testing.T) {
		tests := []string{
			"x + 1",
			"q - 4",
			"a - b + 3",
			"a + b + 3",
			"1 * 2",
			"x + 3 * g",
			"a + b / 2 - c",
			"-a",
			"-2 * a - b * -r",
			"a.b.c",
			"x.y + z.q",
			"\"s\"",
			"\"string tied in a knot\"",
			"'c'",

			"(x + 5)",
			"((x + 5) / 3)",
			"()",

			"{}",
			"{x: int && 5}",
			"{x: int && 5, y: string && \"s\"}",
			"{x: int ?: 5, y: string ?: \"s\"}",
			"{1, 2, 3, 4, 5}",

			"[]",
			"[1, 2, 3, 4, 5]",

			"true and false",
			"a and b",
			"a and b or c",
			"a and not b",
			"not a or b",

			"1 == 2",
			"1 + 1 == 2 / 1",
			"1 + 1 < 2 / 1",
			"1 + 1 <= 2 / 1",
			"1 + 1 >= 2 / 1",

			"x =~ y",
			"x !~ y",

			"int?",
			"float | int?",
			"float & 7.0",

			"f(x: 0)",
			"{a: f(x: 0)}",

			"1..9",
			"x in 1..9",

			"x is Widget",

			"1 when n == 0 | n * f(n - 1) when n > 0",
			"f: (n: int) -> int = 1 when n == 0 | n * f(n - 1) when n > 0",
			"f: (n: int, m: int) -> int = m when n == 0 | n * f(n - 1) when n > 0",

			"x = y + z where {y: 3, z: 5}",
		}

		for _, test := range tests {
			check(test)
		}
	})

}

//---------------------------------------------------------------------------------------------------------------------
