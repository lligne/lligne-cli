//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package codegeneration

import (
	"lligne-cli/internal/lligne/code/model"
	"lligne-cli/internal/lligne/runtime/bytecode"
	"strconv"
)

//=====================================================================================================================

func GenerateByteCode(expression model.ILligneExpression) *bytecode.CodeBlock {
	result := &bytecode.CodeBlock{}

	buildCodeBlock(result, expression)

	result.Stop()

	return result
}

//=====================================================================================================================

func buildCodeBlock(codeBlock *bytecode.CodeBlock, expression model.ILligneExpression) {

	switch expression.TypeCode() {

	case model.ExprTypeBooleanLiteral:
		expr := expression.(*model.LligneBooleanLiteralExpr)
		if expr.Value {
			codeBlock.BoolLoadTrue()
		} else {
			codeBlock.BoolLoadFalse()
		}

	case model.ExprTypeInfixOperation:
		expr := expression.(*model.LligneInfixOperationExpr)
		buildCodeBlock(codeBlock, expr.Operands[0])
		for _, operand := range expr.Operands[1:] {
			buildCodeBlock(codeBlock, operand)
			switch expr.Operator {
			case model.InfixOperatorAdd:
				codeBlock.Int64Add()
			case model.InfixOperatorDivide:
				codeBlock.Int64Divide()
			case model.InfixOperatorEquals:
				codeBlock.Int64Equals()
			case model.InfixOperatorGreaterThan:
				codeBlock.Int64GreaterThan()
			case model.InfixOperatorGreaterThanOrEquals:
				codeBlock.Int64GreaterThanOrEquals()
			case model.InfixOperatorLessThan:
				codeBlock.Int64LessThan()
			case model.InfixOperatorLessThanOrEquals:
				codeBlock.Int64LessThanOrEquals()
			case model.InfixOperatorLogicAnd:
				codeBlock.BoolAnd()
			case model.InfixOperatorLogicOr:
				codeBlock.BoolOr()
			case model.InfixOperatorMultiply:
				codeBlock.Int64Multiply()
			case model.InfixOperatorSubtract:
				codeBlock.Int64Subtract()
			default:
				panic("Unhandled infix operation: " + strconv.Itoa(int(expr.Operator)))
			}
		}

	case model.ExprTypeIntegerLiteral:
		expr := expression.(*model.LligneIntegerLiteralExpr)
		value, _ := strconv.Atoi(expr.Text)
		switch value {
		case 0:
			codeBlock.Int64LoadZero()
		case 1:
			codeBlock.Int64LoadOne()
		default:
			codeBlock.Int64LoadInt16(int16(value))
		}

	case model.ExprTypeParenthesized:
		expr := expression.(*model.LligneParenthesizedExpr)
		if len(expr.Items) == 1 {
			buildCodeBlock(codeBlock, expr.Items[0])
		} else {
			panic("Records not yet handled")
		}

	case model.ExprTypePrefixOperation:
		expr := expression.(*model.LlignePrefixOperationExpr)
		buildCodeBlock(codeBlock, expr.Operand)
		switch expr.Operator {
		case model.PrefixOperatorLogicalNot:
			codeBlock.BoolNot()
		case model.PrefixOperatorNegation:
			codeBlock.Int64Negate()
		default:
			panic("Unhandled prefix operation: " + strconv.Itoa(int(expr.Operator)))
		}

	default:
		panic("Unhandled expression type: " + strconv.Itoa(int(expression.TypeCode())))

	}

}

//=====================================================================================================================
