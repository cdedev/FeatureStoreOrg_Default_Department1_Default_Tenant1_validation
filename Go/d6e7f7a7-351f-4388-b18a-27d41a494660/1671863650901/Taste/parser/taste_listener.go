// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671863650901/Taste.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Taste

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TasteListener is a complete listener for a parse tree produced by TasteParser.
type TasteListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTaste is called when entering the taste production.
	EnterTaste(c *TasteContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTaste is called when exiting the taste production.
	ExitTaste(c *TasteContext)
}
