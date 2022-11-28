// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669648273197/Details.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Details

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DetailsListener is a complete listener for a parse tree produced by DetailsParser.
type DetailsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDetails is called when entering the details production.
	EnterDetails(c *DetailsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDetails is called when exiting the details production.
	ExitDetails(c *DetailsContext)
}
