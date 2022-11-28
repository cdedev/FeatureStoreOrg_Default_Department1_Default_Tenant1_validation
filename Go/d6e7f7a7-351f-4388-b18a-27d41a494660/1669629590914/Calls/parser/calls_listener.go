// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669629590914/Calls.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Calls

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CallsListener is a complete listener for a parse tree produced by CallsParser.
type CallsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCalls is called when entering the calls production.
	EnterCalls(c *CallsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCalls is called when exiting the calls production.
	ExitCalls(c *CallsContext)
}
