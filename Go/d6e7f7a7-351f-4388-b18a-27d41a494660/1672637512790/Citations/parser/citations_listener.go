// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637512790/Citations.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Citations

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CitationsListener is a complete listener for a parse tree produced by CitationsParser.
type CitationsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCitations is called when entering the citations production.
	EnterCitations(c *CitationsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCitations is called when exiting the citations production.
	ExitCitations(c *CitationsContext)
}
