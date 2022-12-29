// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284909159/Nudity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nudity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NudityListener is a complete listener for a parse tree produced by NudityParser.
type NudityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNudity is called when entering the nudity production.
	EnterNudity(c *NudityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNudity is called when exiting the nudity production.
	ExitNudity(c *NudityContext)
}
