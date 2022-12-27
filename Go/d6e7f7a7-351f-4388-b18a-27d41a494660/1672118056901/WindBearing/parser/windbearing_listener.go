// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118056901/WindBearing.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WindBearing

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WindBearingListener is a complete listener for a parse tree produced by WindBearingParser.
type WindBearingListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWindbearing is called when entering the windbearing production.
	EnterWindbearing(c *WindbearingContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWindbearing is called when exiting the windbearing production.
	ExitWindbearing(c *WindbearingContext)
}
