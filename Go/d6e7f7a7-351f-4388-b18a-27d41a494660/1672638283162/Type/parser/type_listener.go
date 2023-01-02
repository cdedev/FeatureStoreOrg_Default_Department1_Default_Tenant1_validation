// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672638283162/Type.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Type

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TypeListener is a complete listener for a parse tree produced by TypeParser.
type TypeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)
}
