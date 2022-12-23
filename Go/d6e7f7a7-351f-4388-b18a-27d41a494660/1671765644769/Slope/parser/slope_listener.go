// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671765644769/Slope.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Slope

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SlopeListener is a complete listener for a parse tree produced by SlopeParser.
type SlopeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSlope is called when entering the slope production.
	EnterSlope(c *SlopeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSlope is called when exiting the slope production.
	ExitSlope(c *SlopeContext)
}
