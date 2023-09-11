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
	"lligne-cli/internal/lligne/code/analysis/pooling"
	"lligne-cli/internal/lligne/code/analysis/structuring"
	"lligne-cli/internal/lligne/code/analysis/typechecking"
	"lligne-cli/internal/lligne/code/codegeneration"
	"lligne-cli/internal/lligne/code/formatting"
	"lligne-cli/internal/lligne/code/parsing"
	"lligne-cli/internal/lligne/code/scanning"
	"lligne-cli/internal/lligne/code/scanning/tokenfilters"
	"lligne-cli/internal/lligne/runtime/bytecode"
	"strings"
	"testing"
)

//---------------------------------------------------------------------------------------------------------------------

func TestExpressionEvaluation(t *testing.T) {

	checkBool := func(sourceCode string) {
		scanOutcome := scanning.Scan(sourceCode)
		scanOutcome = tokenfilters.RemoveDocumentation(scanOutcome)
		parseOutcome := parsing.ParseExpression(scanOutcome)

		assert.Equal(t, sourceCode, formatting.FormatCode(parseOutcome))

		poolOutcome := pooling.PoolConstants(parseOutcome)
		structureOutcome := structuring.StructureRecords(poolOutcome)
		typeCheckOutcome := typechecking.CheckTypes(structureOutcome)
		codeGenOutcome := codegeneration.GenerateByteCode(typeCheckOutcome)

		//print(codeBlock.Disassemble())

		stringPool := codeGenOutcome.StringConstants.Clone()
		typePool := codeGenOutcome.TypeConstants.Clone()

		interpreter := bytecode.NewInterpreter(codeGenOutcome.CodeBlock, stringPool, typePool)
		machine := bytecode.NewMachine()

		interpreter.Execute(machine)

		actual := machine.BoolGetResult()

		assert.True(t, actual, "For source code: "+sourceCode)
	}

	checkSampleFile := func(t *testing.T, sampleContent string) {

		samples := strings.Split(sampleContent, "•")

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
		checkSampleFile(t, sample4)
		checkSampleFile(t, sample5)
		checkSampleFile(t, sample6)
		checkSampleFile(t, sample7)

	})

}

//---------------------------------------------------------------------------------------------------------------------

//go:embed bool/logic-true.lligne-tests
var sample1 string

//go:embed float64/float64-comparisons-true.lligne-tests
var sample2 string

//go:embed int64/int64-comparisons-true.lligne-tests
var sample3 string

//go:embed record/record-comparisons.lligne-tests
var sample4 string

//go:embed string/string-comparisons.lligne-tests
var sample5 string

//go:embed string/string-concatenation.lligne-tests
var sample6 string

//go:embed types/built-in-types.lligne-tests
var sample7 string

//---------------------------------------------------------------------------------------------------------------------
