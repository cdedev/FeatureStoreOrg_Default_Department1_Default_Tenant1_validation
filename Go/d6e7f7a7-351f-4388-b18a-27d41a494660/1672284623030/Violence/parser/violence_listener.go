// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284623030/Violence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Violence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ViolenceListener is a complete listener for a parse tree produced by ViolenceParser.
type ViolenceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterViolence is called when entering the violence production.
	EnterViolence(c *ViolenceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitViolence is called when exiting the violence production.
	ExitViolence(c *ViolenceContext)
}
