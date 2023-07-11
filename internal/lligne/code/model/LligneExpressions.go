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
	ExprTypeIntegerLiteral
	ExprTypeLeadingDocumentation
	ExprTypeMultilineStringLiteral
	ExprTypeStringLiteral
)

//---------------------------------------------------------------------------------------------------------------------

// LligneIdentifierExpr represents a single identifier.
type LligneIdentifierExpr struct {
	SourcePos int
	Name      string
}

func (e *LligneIdentifierExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.SourcePos)
}

func (e *LligneIdentifierExpr) SExpression() string {
	return "(identifier " + e.Name + ")"
}

func (e *LligneIdentifierExpr) TypeCode() LligneExprType {
	return ExprTypeIdentifier
}

//---------------------------------------------------------------------------------------------------------------------

// LligneIntegerLiteralExpr represents a single integer literal.
type LligneIntegerLiteralExpr struct {
	SourcePos int
	Text      string
}

func (e *LligneIntegerLiteralExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.SourcePos)
}

func (e *LligneIntegerLiteralExpr) SExpression() string {
	return "(int " + e.Text + ")"
}

func (e *LligneIntegerLiteralExpr) TypeCode() LligneExprType {
	return ExprTypeIntegerLiteral
}

//---------------------------------------------------------------------------------------------------------------------

// LligneLeadingDocumentationExpr represents lines of leading documentation.
type LligneLeadingDocumentationExpr struct {
	SourcePos int
	Text      string
}

func (e *LligneLeadingDocumentationExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.SourcePos)
}

func (e *LligneLeadingDocumentationExpr) SExpression() string {
	return "(leadingdoc\n" + e.Text + ")"
}

func (e *LligneLeadingDocumentationExpr) TypeCode() LligneExprType {
	return ExprTypeLeadingDocumentation
}

//---------------------------------------------------------------------------------------------------------------------

// LligneMultilineStringLiteralExpr represents a single integer literal.
type LligneMultilineStringLiteralExpr struct {
	SourcePos int
	Text      string
}

func (e *LligneMultilineStringLiteralExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.SourcePos)
}

func (e *LligneMultilineStringLiteralExpr) SExpression() string {
	return "(multilinestr\n" + e.Text + ")"
}

func (e *LligneMultilineStringLiteralExpr) TypeCode() LligneExprType {
	return ExprTypeMultilineStringLiteral
}

//---------------------------------------------------------------------------------------------------------------------

// LligneStringLiteralExpr represents a single integer literal.
type LligneStringLiteralExpr struct {
	SourcePos int
	Text      string
}

func (e *LligneStringLiteralExpr) GetOrigin(tracker scanning.ILligneTokenOriginTracker) scanning.LligneOrigin {
	return tracker.GetOrigin(e.SourcePos)
}

func (e *LligneStringLiteralExpr) SExpression() string {
	return "(string " + e.Text + ")"
}

func (e *LligneStringLiteralExpr) TypeCode() LligneExprType {
	return ExprTypeStringLiteral
}

//---------------------------------------------------------------------------------------------------------------------
