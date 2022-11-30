// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669795091007/Carbon.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carbon

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CarbonListener is a complete listener for a parse tree produced by CarbonParser.
type CarbonListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCarbon is called when entering the carbon production.
	EnterCarbon(c *CarbonContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCarbon is called when exiting the carbon production.
	ExitCarbon(c *CarbonContext)
}
