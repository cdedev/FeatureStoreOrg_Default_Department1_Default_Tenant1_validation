// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672204772426/Lightbill.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lightbill

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LightbillListener is a complete listener for a parse tree produced by LightbillParser.
type LightbillListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLightbill is called when entering the lightbill production.
	EnterLightbill(c *LightbillContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLightbill is called when exiting the lightbill production.
	ExitLightbill(c *LightbillContext)
}
