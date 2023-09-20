//
// # Tests of the parser for Lligne code
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package codegeneration

import (
	"github.com/stretchr/testify/assert"
	"lligne-cli/internal/lligne/code/analysis/nameresolution"
	"lligne-cli/internal/lligne/code/analysis/pooling"
	"lligne-cli/internal/lligne/code/analysis/structuring"
	"lligne-cli/internal/lligne/code/analysis/typechecking"
	"lligne-cli/internal/lligne/code/parsing"
	"lligne-cli/internal/lligne/code/scanning"
	"lligne-cli/internal/lligne/code/scanning/tokenfilters"
	"lligne-cli/internal/lligne/runtime/bytecode"
	"lligne-cli/internal/lligne/runtime/pools"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func runInterpreter(sourceCode string) (*bytecode.Machine, *pools.StringPool) {
	scanOutcome := scanning.Scan(sourceCode)
	scanOutcome = tokenfilters.RemoveDocumentation(scanOutcome)
	parseOutcome := parsing.ParseExpression(scanOutcome)
	poolOutcome := pooling.PoolConstants(parseOutcome)
	structureOutcome := structuring.StructureRecords(poolOutcome)
	resolutionOutcome := nameresolution.ResolveNames(structureOutcome)
	typeCheckOutcome := typechecking.CheckTypes(resolutionOutcome)
	codeGenOutcome := GenerateByteCode(typeCheckOutcome)

	//print(codeBlock.Disassemble())

	stringPool := codeGenOutcome.StringConstants.Clone()
	typePool := codeGenOutcome.TypeConstants.Clone()

	interpreter := bytecode.NewInterpreter(codeGenOutcome.CodeBlock, stringPool, typePool)
	machine := bytecode.NewMachine()

	interpreter.Execute(machine)

	return machine, stringPool
}

//---------------------------------------------------------------------------------------------------------------------

func TestGenerateBoolByteCode(t *testing.T) {

	checkBool := func(sourceCode string, expected bool) {
		machine, _ := runInterpreter(sourceCode)

		actual := machine.BoolGetResult()

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
		machine, _ := runInterpreter(sourceCode)

		actual := machine.Int64GetResult()

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
		machine, _ := runInterpreter(sourceCode)

		actual := machine.Float64GetResult()

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
		machine, stringPool := runInterpreter(sourceCode)

		actual := machine.StringGetResult(stringPool)

		assert.Equal(t, expected, actual, "For source code: "+sourceCode)
	}

	t.Run("String expression evaluations", func(t *testing.T) {
		type exprOutcome struct {
			sourceCode    string
			expectedValue string
		}

		tests := []exprOutcome{
			{`'A string'`, "A string"},
			{`'one' + 'two'`, "onetwo"},
			{`'one' + ('two' + 'three')`, "onetwothree"},
			{`"A string"`, "A string"},
			{`"one" + "two"`, "onetwo"},
			{`"one" + ("two" + "three")`, "onetwothree"},
		}
		for _, test := range tests {
			checkString(test.sourceCode, test.expectedValue)
		}
	})

}

//---------------------------------------------------------------------------------------------------------------------
