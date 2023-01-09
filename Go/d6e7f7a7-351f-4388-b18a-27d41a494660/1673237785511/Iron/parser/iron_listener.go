// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673237785511/Iron.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Iron

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IronListener is a complete listener for a parse tree produced by IronParser.
type IronListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIron is called when entering the iron production.
	EnterIron(c *IronContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIron is called when exiting the iron production.
	ExitIron(c *IronContext)
}
