// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674795991425/Hispanic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hispanic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HispanicListener is a complete listener for a parse tree produced by HispanicParser.
type HispanicListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHispanic is called when entering the hispanic production.
	EnterHispanic(c *HispanicContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHispanic is called when exiting the hispanic production.
	ExitHispanic(c *HispanicContext)
}
