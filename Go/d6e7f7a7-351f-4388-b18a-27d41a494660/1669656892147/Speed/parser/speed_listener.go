// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656892147/Speed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Speed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SpeedListener is a complete listener for a parse tree produced by SpeedParser.
type SpeedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSpeed is called when entering the speed production.
	EnterSpeed(c *SpeedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSpeed is called when exiting the speed production.
	ExitSpeed(c *SpeedContext)
}
