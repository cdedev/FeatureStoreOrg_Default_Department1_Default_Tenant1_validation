// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675758462392/Clicks.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Clicks

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ClicksListener is a complete listener for a parse tree produced by ClicksParser.
type ClicksListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterClicks is called when entering the clicks production.
	EnterClicks(c *ClicksContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitClicks is called when exiting the clicks production.
	ExitClicks(c *ClicksContext)
}
