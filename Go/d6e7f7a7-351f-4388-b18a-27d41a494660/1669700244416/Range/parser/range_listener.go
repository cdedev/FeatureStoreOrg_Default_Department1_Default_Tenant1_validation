// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700244416/Range.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Range

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RangeListener is a complete listener for a parse tree produced by RangeParser.
type RangeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRange is called when entering the range production.
	EnterRange(c *RangeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRange is called when exiting the range production.
	ExitRange(c *RangeContext)
}
