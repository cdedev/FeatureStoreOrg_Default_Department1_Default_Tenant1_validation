// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669795209848/Carbondioxide.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carbondioxide

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CarbondioxideListener is a complete listener for a parse tree produced by CarbondioxideParser.
type CarbondioxideListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCarbondioxide is called when entering the carbondioxide production.
	EnterCarbondioxide(c *CarbondioxideContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCarbondioxide is called when exiting the carbondioxide production.
	ExitCarbondioxide(c *CarbondioxideContext)
}
