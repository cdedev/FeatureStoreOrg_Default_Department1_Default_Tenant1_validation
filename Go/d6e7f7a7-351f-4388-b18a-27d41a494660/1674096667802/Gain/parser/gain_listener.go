// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674096667802/Gain.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gain

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GainListener is a complete listener for a parse tree produced by GainParser.
type GainListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGain is called when entering the gain production.
	EnterGain(c *GainContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGain is called when exiting the gain production.
	ExitGain(c *GainContext)
}
