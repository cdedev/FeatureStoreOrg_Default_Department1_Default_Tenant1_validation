// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671862884154/Decrement.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Decrement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DecrementListener is a complete listener for a parse tree produced by DecrementParser.
type DecrementListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDecrement is called when entering the decrement production.
	EnterDecrement(c *DecrementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDecrement is called when exiting the decrement production.
	ExitDecrement(c *DecrementContext)
}
