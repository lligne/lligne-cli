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

	switch expression.ExprType() {

	case model.ExprTypeBooleanLiteral:
		expr := expression.(*model.LligneBooleanLiteralExpr)
		if expr.Value {
			codeBlock.BoolLoadTrue()
		} else {
			codeBlock.BoolLoadFalse()
		}

	case model.ExprTypeFloatingPointLiteral:
		expr := expression.(*model.LligneFloatingPointLiteralExpr)
		value, _ := strconv.ParseFloat(expr.Text, 64)
		switch value {
		case 0:
			codeBlock.Float64LoadZero()
		case 1:
			codeBlock.Float64LoadOne()
		default:
			codeBlock.Float64LoadFloat64(value)
		}

	case model.ExprTypeInfixOperation:
		expr := expression.(*model.LligneInfixOperationExpr)
		buildCodeBlock(codeBlock, expr.Operands[0])
		for _, operand := range expr.Operands[1:] {
			buildCodeBlock(codeBlock, operand)
			switch expr.Operator {
			case model.InfixOperatorAdd:
				if expr.TypeInfo().BaseType() == model.BaseTypeInt64 {
					codeBlock.Int64Add()
				} else {
					codeBlock.Float64Add()
				}
			case model.InfixOperatorDivide:
				if expr.TypeInfo().BaseType() == model.BaseTypeInt64 {
					codeBlock.Int64Divide()
				} else {
					codeBlock.Float64Divide()
				}
			case model.InfixOperatorEquals:
				if expr.TypeInfo().BaseType() == model.BaseTypeInt64 {
					codeBlock.Int64Equals()
				} else {
					codeBlock.Float64Equals()
				}
			case model.InfixOperatorGreaterThan:
				if expr.TypeInfo().BaseType() == model.BaseTypeInt64 {
					codeBlock.Int64GreaterThan()
				} else {
					codeBlock.Float64GreaterThan()
				}
			case model.InfixOperatorGreaterThanOrEquals:
				if expr.TypeInfo().BaseType() == model.BaseTypeInt64 {
					codeBlock.Int64GreaterThanOrEquals()
				} else {
					codeBlock.Float64GreaterThanOrEquals()
				}
			case model.InfixOperatorLessThan:
				if expr.TypeInfo().BaseType() == model.BaseTypeInt64 {
					codeBlock.Int64LessThan()
				} else {
					codeBlock.Float64LessThan()
				}
			case model.InfixOperatorLessThanOrEquals:
				if expr.TypeInfo().BaseType() == model.BaseTypeInt64 {
					codeBlock.Int64LessThanOrEquals()
				} else {
					codeBlock.Float64LessThanOrEquals()
				}
			case model.InfixOperatorLogicAnd:
				codeBlock.BoolAnd()
			case model.InfixOperatorLogicOr:
				codeBlock.BoolOr()
			case model.InfixOperatorMultiply:
				if expr.TypeInfo().BaseType() == model.BaseTypeInt64 {
					codeBlock.Int64Multiply()
				} else {
					codeBlock.Float64Multiply()
				}
			case model.InfixOperatorSubtract:
				if expr.TypeInfo().BaseType() == model.BaseTypeInt64 {
					codeBlock.Int64Subtract()
				} else {
					codeBlock.Float64Subtract()
				}
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
			if expr.TypeInfo().BaseType() == model.BaseTypeInt64 {
				codeBlock.Int64Negate()
			} else {
				codeBlock.Float64Negate()
			}
		default:
			panic("Unhandled prefix operation: " + strconv.Itoa(int(expr.Operator)))
		}

	default:
		panic("Unhandled expression type: " + strconv.Itoa(int(expression.ExprType())))

	}

}

//=====================================================================================================================
