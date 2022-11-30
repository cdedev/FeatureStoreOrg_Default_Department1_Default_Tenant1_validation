// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669792843997/Width.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Width

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WidthListener is a complete listener for a parse tree produced by WidthParser.
type WidthListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWidth is called when entering the width production.
	EnterWidth(c *WidthContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWidth is called when exiting the width production.
	ExitWidth(c *WidthContext)
}
