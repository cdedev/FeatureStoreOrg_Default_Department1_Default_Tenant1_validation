// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671601452315/SourceID.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SourceID

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SourceIDListener is a complete listener for a parse tree produced by SourceIDParser.
type SourceIDListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSourceid is called when entering the sourceid production.
	EnterSourceid(c *SourceidContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSourceid is called when exiting the sourceid production.
	ExitSourceid(c *SourceidContext)
}
