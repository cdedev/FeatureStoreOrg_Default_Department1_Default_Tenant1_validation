// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675232763379/Ratio.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ratio

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RatioListener is a complete listener for a parse tree produced by RatioParser.
type RatioListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRatio is called when entering the ratio production.
	EnterRatio(c *RatioContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRatio is called when exiting the ratio production.
	ExitRatio(c *RatioContext)
}
