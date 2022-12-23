// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671767294086/Killed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Killed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// KilledListener is a complete listener for a parse tree produced by KilledParser.
type KilledListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterKilled is called when entering the killed production.
	EnterKilled(c *KilledContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitKilled is called when exiting the killed production.
	ExitKilled(c *KilledContext)
}
