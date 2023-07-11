//
// (C) Copyright 2023 Martin E. Nordberg III
// Apache 2.0 License
//

package model

import "lligne-cli/internal/lligne/code/scanning"

//---------------------------------------------------------------------------------------------------------------------

// ILligneExpression is the interface to an expression AST node.
type ILligneExpression interface {
	GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin
	SExpression() string
	TypeCode() LligneExprType
}

//---------------------------------------------------------------------------------------------------------------------

// LligneExprType is an enumeration of Lligne expression types.
type LligneExprType int

const (
	ExprTypeIdentifier LligneExprType = iota
)

//---------------------------------------------------------------------------------------------------------------------

// lligneIdentifierExpr represents a single identifier.
type lligneIdentifierExpr struct {
	sourcePos int
	name      string
}

func NewLligneIdentifierExpr(sourcePos int, name string) ILligneExpression {
	return &lligneIdentifierExpr{
		sourcePos: sourcePos,
		name:      name,
	}
}

func (e *lligneIdentifierExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.sourcePos)
}

func (e *lligneIdentifierExpr) SExpression() string {
	return "(identifier " + e.name + ")"
}

func (e *lligneIdentifierExpr) TypeCode() LligneExprType {
	return ExprTypeIdentifier
}

//---------------------------------------------------------------------------------------------------------------------
