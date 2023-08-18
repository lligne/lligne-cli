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
		tokens, _ := scanning.Scan(sourceCode)

		tokens = scanning.RemoveDocumentation(tokens)

		model := parsing.ParseExpression(sourceCode, tokens)

		typedModel := typechecking.TypeCheckExpr(model)

		codeBlock := GenerateByteCode(typedModel)

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
			{"false and false", false},
			{"true or false", true},
			{"true or true", true},
			{"false or false", false},
			{"not true", false},
			{"not false", true},
			{"true and not false", true},

			{"2 == 1 + 1", true},
			{"3 == 1 + 1", false},
			{"2 <= 1 + 1", true},
			{"3 <= 1 + 1", false},
			{"2 >= 1 + 1", true},
			{"1 >= 1 + 1", false},
			{"1 <= 1 + 1", true},
			{"2 <= 1 + 1", true},
			{"3 <= 1 + 1", false},
			{"3 >= 1 + 1", true},
			{"2 >= 1 + 1", true},
			{"1 >= 1 + 1", false},
			{"1 < 1 + 1", true},
			{"2 < 1 + 1", false},
			{"3 > 1 + 1", true},
			{"2 > 1 + 1", false},

			{"2.0 == 1.0 + 1.0", true},
			{"3.0 == 1.0 + 1.0", false},
			{"2.0 <= 1.0 + 1.0", true},
			{"3.0 <= 1.0 + 1.0", false},
			{"2.0 >= 1.0 + 1.0", true},
			{"1.0 >= 1.0 + 1.0", false},
			{"1.0 <= 1.0 + 1.0", true},
			{"2.0 <= 1.0 + 1.0", true},
			{"3.0 <= 1.0 + 1.0", false},
			{"3.0 >= 1.0 + 1.0", true},
			{"2.0 >= 1.0 + 1.0", true},
			{"1.0 >= 1.0 + 1.0", false},
			{"1.0 < 1.0 + 1.0", true},
			{"2.0 < 1.0 + 1.0", false},
			{"3.0 > 1.0 + 1.0", true},
			{"2.0 > 1.0 + 1.0", false},
		}
		for _, test := range tests {
			checkBool(test.sourceCode, test.expectedValue)
		}
	})

}

//---------------------------------------------------------------------------------------------------------------------

func TestGenerateInt64ByteCode(t *testing.T) {

	checkInt64 := func(sourceCode string, expected int64) {
		tokens, _ := scanning.Scan(sourceCode)

		model := parsing.ParseExpression(sourceCode, tokens)

		typedModel := typechecking.TypeCheckExpr(model)

		codeBlock := GenerateByteCode(typedModel)

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
			{"1 + 1", 2},
			{"1 + 2", 3},
			{"2 + 1", 3},
			{"1 + 2 + 7", 10},
			{"(1 + 2) + (7 + 5)", 15},
			{"20 - 1", 19},
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
		tokens, _ := scanning.Scan(sourceCode)

		model := parsing.ParseExpression(sourceCode, tokens)

		typedModel := typechecking.TypeCheckExpr(model)

		codeBlock := GenerateByteCode(typedModel)

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

func TestGenerateStringByteCode(t *testing.T) {

	checkString := func(sourceCode string, expected string) {
		tokens, _ := scanning.Scan(sourceCode)

		model := parsing.ParseExpression(sourceCode, tokens)

		typedModel := typechecking.TypeCheckExpr(model)

		codeBlock := GenerateByteCode(typedModel)

		//print(codeBlock.Disassemble())

		interpreter := &bytecode.Interpreter{}
		machine := bytecode.NewMachine()

		interpreter.Execute(machine, codeBlock)

		actual := interpreter.StringGetResult(machine, codeBlock)

		assert.Equal(t, expected, actual, "For source code: "+sourceCode)
	}

	t.Run("String expression evaluations", func(t *testing.T) {
		type exprOutcome struct {
			sourceCode    string
			expectedValue string
		}

		tests := []exprOutcome{
			//{"'A string'", "A string"},
			{"'one' + 'two'", "onetwo"},
		}
		for _, test := range tests {
			checkString(test.sourceCode, test.expectedValue)
		}
	})

}

//---------------------------------------------------------------------------------------------------------------------
