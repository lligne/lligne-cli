//
// # Tests of the parser for Lligne code
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package codegeneration

import (
	"github.com/stretchr/testify/assert"
	"lligne-cli/internal/lligne/code/parsing"
	"lligne-cli/internal/lligne/code/scanning"
	"lligne-cli/internal/lligne/runtime/bytecode"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestGenerateByteCode(t *testing.T) {

	checkInt64 := func(sourceCode string, expected int64) {
		scanner := scanning.NewLligneBufferedScanner(
			scanning.NewLligneDocumentationHandlingScanner(
				sourceCode,
				scanning.NewLligneScanner(sourceCode),
			),
		)
		parser := parsing.NewLligneParser(scanner)
		model := parser.ParseExpression()

		codeBlock := GenerateByteCode(model)

		//disassembler := &bytecode.Disassembler{}
		//codeBlock.Execute(disassembler)
		//print(disassembler.GetOutput())

		interpreter := &bytecode.Interpreter{}

		codeBlock.Execute(interpreter)

		actual := interpreter.Int64GetResult()

		assert.Equal(t, expected, actual, "For source code: "+sourceCode)
	}

	t.Run("Int64 expression evaluations", func(t *testing.T) {
		type exprOutcome struct {
			sourceCode    string
			expectedValue int64
		}

		tests := []exprOutcome{
			{"0 + 1", 1},
			{"1 + 2", 3},
			{"1 + 2 + 7", 10},
			{"(1 + 2) + (7 + 5)", 15},
			{"20 - 2", 18},
			{"20 - 2 - 4", 14},
			{"(1 + 2) + (7 - 5)", 5},
			{"(22 + 2) - (7 - 5)", 22},
			{"20 * 2", 40},
			{"(5 + 6 - 1) * (0 + 1 + 2 + 3)", 60},
			{"20 / 2", 10},
			{"20 / (1 + 1)", 10},
			{"-7", -7},
			{"-(7 - 3) + 1", -3},
		}
		for _, test := range tests {
			checkInt64(test.sourceCode, test.expectedValue)
		}
	})
}

//---------------------------------------------------------------------------------------------------------------------