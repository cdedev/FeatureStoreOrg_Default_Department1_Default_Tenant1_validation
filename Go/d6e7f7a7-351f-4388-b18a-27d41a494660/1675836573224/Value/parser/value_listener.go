// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675836573224/Value.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Value

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ValueListener is a complete listener for a parse tree produced by ValueParser.
type ValueListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
