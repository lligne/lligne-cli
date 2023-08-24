//
// # Tests of the parser for Lligne code
//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package tests

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"lligne-cli/internal/lligne/code/codegeneration"
	"lligne-cli/internal/lligne/code/formatting"
	"lligne-cli/internal/lligne/code/parsing"
	"lligne-cli/internal/lligne/code/scanning"
	"lligne-cli/internal/lligne/code/typechecking"
	"lligne-cli/internal/lligne/runtime/bytecode"
	"strings"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestExpressionEvaluation(t *testing.T) {

	checkBool := func(sourceCode string) {
		tokens, _ := scanning.Scan(sourceCode)

		tokens = scanning.RemoveDocumentation(tokens)

		expression := parsing.ParseExpression(sourceCode, tokens)

		assert.Equal(t, sourceCode, formatting.FormatExpr(sourceCode, expression))

		typedModel := typechecking.TypeCheckExpr(sourceCode, expression)

		codeBlock := codegeneration.GenerateByteCode(typedModel)

		//print(codeBlock.Disassemble())

		interpreter := &bytecode.Interpreter{}
		machine := bytecode.NewMachine()

		interpreter.Execute(machine, codeBlock)

		actual := interpreter.BoolGetResult(machine)

		assert.True(t, actual, "For source code: "+sourceCode)
	}

	checkSampleFile := func(t *testing.T, sampleContent string) {

		samples := strings.Split(sampleContent, "\n")

		for _, sample := range samples {
			expression := strings.TrimSpace(sample)
			if len(expression) > 0 {
				checkBool(expression)
			}
		}

	}

	t.Run("Boolean expression evaluations", func(t *testing.T) {

		checkSampleFile(t, sample1)
		checkSampleFile(t, sample2)
		checkSampleFile(t, sample3)

	})

}

//---------------------------------------------------------------------------------------------------------------------

//go:embed float64/float64-comparisons-true.lligne-tests
var sample1 string

//go:embed int64/int64-comparisons-true.lligne-tests
var sample2 string

//go:embed bool/logic-true.lligne-tests
var sample3 string

//---------------------------------------------------------------------------------------------------------------------
