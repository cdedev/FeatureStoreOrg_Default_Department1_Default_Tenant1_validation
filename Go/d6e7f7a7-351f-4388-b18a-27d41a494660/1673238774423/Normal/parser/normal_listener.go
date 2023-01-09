// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238774423/Normal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Normal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NormalListener is a complete listener for a parse tree produced by NormalParser.
type NormalListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNormal is called when entering the normal production.
	EnterNormal(c *NormalContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNormal is called when exiting the normal production.
	ExitNormal(c *NormalContext)
}
