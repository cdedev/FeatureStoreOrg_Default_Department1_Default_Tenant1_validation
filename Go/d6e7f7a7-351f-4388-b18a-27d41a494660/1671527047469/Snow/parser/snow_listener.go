// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671527047469/Snow.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Snow

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SnowListener is a complete listener for a parse tree produced by SnowParser.
type SnowListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSnow is called when entering the snow production.
	EnterSnow(c *SnowContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSnow is called when exiting the snow production.
	ExitSnow(c *SnowContext)
}
