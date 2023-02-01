// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675232880958/Skew.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Skew

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SkewListener is a complete listener for a parse tree produced by SkewParser.
type SkewListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSkew is called when entering the skew production.
	EnterSkew(c *SkewContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSkew is called when exiting the skew production.
	ExitSkew(c *SkewContext)
}
