// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118879751/MaxSpeed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MaxSpeed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MaxSpeedListener is a complete listener for a parse tree produced by MaxSpeedParser.
type MaxSpeedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMaxspeed is called when entering the maxspeed production.
	EnterMaxspeed(c *MaxspeedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMaxspeed is called when exiting the maxspeed production.
	ExitMaxspeed(c *MaxspeedContext)
}
