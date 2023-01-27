// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674792391904/Waterproof.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Waterproof

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WaterproofListener is a complete listener for a parse tree produced by WaterproofParser.
type WaterproofListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWaterproof is called when entering the waterproof production.
	EnterWaterproof(c *WaterproofContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWaterproof is called when exiting the waterproof production.
	ExitWaterproof(c *WaterproofContext)
}
