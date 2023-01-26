// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674721093144/Leave.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Leave

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LeaveListener is a complete listener for a parse tree produced by LeaveParser.
type LeaveListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLeave is called when entering the leave production.
	EnterLeave(c *LeaveContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLeave is called when exiting the leave production.
	ExitLeave(c *LeaveContext)
}
