// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669627420501/Open.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Open

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OpenListener is a complete listener for a parse tree produced by OpenParser.
type OpenListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOpen is called when entering the open production.
	EnterOpen(c *OpenContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOpen is called when exiting the open production.
	ExitOpen(c *OpenContext)
}
