// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671607995826/Power.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Power

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PowerListener is a complete listener for a parse tree produced by PowerParser.
type PowerListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPower is called when entering the power production.
	EnterPower(c *PowerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPower is called when exiting the power production.
	ExitPower(c *PowerContext)
}
