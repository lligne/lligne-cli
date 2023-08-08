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
	"lligne-cli/internal/lligne/code/typechecking"
	"lligne-cli/internal/lligne/runtime/bytecode"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestGenerateBoolByteCode(t *testing.T) {

	checkBool := func(sourceCode string, expected bool) {
		scanner := scanning.NewLligneBufferedScanner(
			scanning.NewLligneDocumentationHandlingScanner(
				sourceCode,
				scanning.NewLligneScanner(sourceCode),
			),
		)
		parser := parsing.NewLligneParser(scanner)
		model := parser.ParseExpression()
		typechecking.DetermineTypes(&model)

		codeBlock := GenerateByteCode(model)

		//print(codeBlock.Disassemble())

		interpreter := &bytecode.Interpreter{}
		machine := bytecode.NewMachine()

		interpreter.Execute(machine, codeBlock)

		actual := interpreter.BoolGetResult(machine)

		assert.Equal(t, expected, actual, "For source code: "+sourceCode)
	}

	t.Run("Boolean expression evaluations", func(t *testing.T) {
		type exprOutcome struct {
			sourceCode    string
			expectedValue bool
		}

		tests := []exprOutcome{
			{"true", true},
			{"false", false},
			{"true and false", false},
			{"true and true", true},
			{"not true", false},
			{"not false", true},
			{"true and not false", true},

			{"2 == 1 + 1", true},
			{"2 <= 1 + 1", true},
			{"2 >= 1 + 1", true},
			{"1 <= 1 + 1", true},
			{"3 >= 1 + 1", true},
			{"1 < 1 + 1", true},
			{"3 > 1 + 1", true},
		}
		for _, test := range tests {
			checkBool(test.sourceCode, test.expectedValue)
		}
	})

}

//---------------------------------------------------------------------------------------------------------------------

func TestGenerateInt64ByteCode(t *testing.T) {

	checkInt64 := func(sourceCode string, expected int64) {
		scanner := scanning.NewLligneBufferedScanner(
			scanning.NewLligneDocumentationHandlingScanner(
				sourceCode,
				scanning.NewLligneScanner(sourceCode),
			),
		)
		parser := parsing.NewLligneParser(scanner)
		model := parser.ParseExpression()
		typechecking.DetermineTypes(&model)

		codeBlock := GenerateByteCode(model)

		//print(codeBlock.Disassemble())

		interpreter := &bytecode.Interpreter{}
		machine := bytecode.NewMachine()

		interpreter.Execute(machine, codeBlock)

		actual := interpreter.Int64GetResult(machine)

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

func TestGenerateFloat64ByteCode(t *testing.T) {

	checkFloat64 := func(sourceCode string, expected float64) {
		scanner := scanning.NewLligneBufferedScanner(
			scanning.NewLligneDocumentationHandlingScanner(
				sourceCode,
				scanning.NewLligneScanner(sourceCode),
			),
		)
		parser := parsing.NewLligneParser(scanner)
		model := parser.ParseExpression()
		typechecking.DetermineTypes(&model)

		codeBlock := GenerateByteCode(model)

		//print(codeBlock.Disassemble())

		interpreter := &bytecode.Interpreter{}
		machine := bytecode.NewMachine()

		interpreter.Execute(machine, codeBlock)

		actual := interpreter.Float64GetResult(machine)

		assert.Equal(t, expected, actual, "For source code: "+sourceCode)
	}

	t.Run("Float64 expression evaluations", func(t *testing.T) {
		type exprOutcome struct {
			sourceCode    string
			expectedValue float64
		}

		tests := []exprOutcome{
			{"0.0 + 1.0", 1.0},
			{"1.0 + 2.0", 3.0},
			{"1.0 + 2.0 + 7.0", 10.0},
			{"(1.0 + 2.0) + (7.0 + 5.0)", 15.0},
			{"20.0 - 2.0", 18.0},
			{"20.0 - 2.0 - 4.0", 14.0},
			{"(1.0 + 2.0) + (7.0 - 5.0)", 5.0},
			{"(22.0 + 2.0) - (7.0 - 5.0)", 22},
			{"20.0 * 2.0", 40.0},
			{"(5.0 + 6.0 - 1.0) * (0.0 + 1.0 + 2.0 + 3.0)", 60.0},
			{"20.0 / 2.0", 10.0},
			{"20.0 / (1.0 + 1.0)", 10.0},
			{"-7.0", -7},
			{"-(7.0 - 3.0) + 1.0", -3.0},
		}
		for _, test := range tests {
			checkFloat64(test.sourceCode, test.expectedValue)
		}
	})

}

//---------------------------------------------------------------------------------------------------------------------
