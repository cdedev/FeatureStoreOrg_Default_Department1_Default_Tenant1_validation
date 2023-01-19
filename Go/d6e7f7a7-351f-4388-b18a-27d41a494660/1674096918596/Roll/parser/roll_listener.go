// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674096918596/Roll.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Roll

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RollListener is a complete listener for a parse tree produced by RollParser.
type RollListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRoll is called when entering the roll production.
	EnterRoll(c *RollContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRoll is called when exiting the roll production.
	ExitRoll(c *RollContext)
}
