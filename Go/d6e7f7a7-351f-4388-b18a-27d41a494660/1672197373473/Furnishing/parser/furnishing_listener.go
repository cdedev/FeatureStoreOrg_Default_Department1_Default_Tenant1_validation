// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672197373473/Furnishing.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Furnishing

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FurnishingListener is a complete listener for a parse tree produced by FurnishingParser.
type FurnishingListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFurnishing is called when entering the furnishing production.
	EnterFurnishing(c *FurnishingContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFurnishing is called when exiting the furnishing production.
	ExitFurnishing(c *FurnishingContext)
}
