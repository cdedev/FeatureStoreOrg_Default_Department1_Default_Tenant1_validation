// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669626531262/Weight.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Weight

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WeightListener is a complete listener for a parse tree produced by WeightParser.
type WeightListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWeight is called when entering the weight production.
	EnterWeight(c *WeightContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWeight is called when exiting the weight production.
	ExitWeight(c *WeightContext)
}
