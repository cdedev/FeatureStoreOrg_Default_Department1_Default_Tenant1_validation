// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673929021260/Status.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Status

import "github.com/antlr/antlr4/runtime/Go/antlr"

// StatusListener is a complete listener for a parse tree produced by StatusParser.
type StatusListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterStatus is called when entering the status production.
	EnterStatus(c *StatusContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitStatus is called when exiting the status production.
	ExitStatus(c *StatusContext)
}
