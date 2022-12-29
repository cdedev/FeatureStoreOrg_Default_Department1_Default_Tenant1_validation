// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672285317914/Detail.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Detail

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DetailListener is a complete listener for a parse tree produced by DetailParser.
type DetailListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDetail is called when entering the detail production.
	EnterDetail(c *DetailContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDetail is called when exiting the detail production.
	ExitDetail(c *DetailContext)
}
