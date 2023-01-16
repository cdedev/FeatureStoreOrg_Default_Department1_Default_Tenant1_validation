// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673857505294/FastChargeSpeed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FastChargeSpeed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FastChargeSpeedListener is a complete listener for a parse tree produced by FastChargeSpeedParser.
type FastChargeSpeedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFastchargespeed is called when entering the fastchargespeed production.
	EnterFastchargespeed(c *FastchargespeedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFastchargespeed is called when exiting the fastchargespeed production.
	ExitFastchargespeed(c *FastchargespeedContext)
}
