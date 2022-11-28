// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669651767664/Source.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Source

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SourceListener is a complete listener for a parse tree produced by SourceParser.
type SourceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSource is called when entering the source production.
	EnterSource(c *SourceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSource is called when exiting the source production.
	ExitSource(c *SourceContext)
}
