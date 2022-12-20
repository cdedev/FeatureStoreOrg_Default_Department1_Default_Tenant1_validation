// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671515756311/Positive.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Positive

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PositiveListener is a complete listener for a parse tree produced by PositiveParser.
type PositiveListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPositive is called when entering the positive production.
	EnterPositive(c *PositiveContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPositive is called when exiting the positive production.
	ExitPositive(c *PositiveContext)
}
