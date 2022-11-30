// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669794328179/Views.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Views

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ViewsListener is a complete listener for a parse tree produced by ViewsParser.
type ViewsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterViews is called when entering the views production.
	EnterViews(c *ViewsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitViews is called when exiting the views production.
	ExitViews(c *ViewsContext)
}
