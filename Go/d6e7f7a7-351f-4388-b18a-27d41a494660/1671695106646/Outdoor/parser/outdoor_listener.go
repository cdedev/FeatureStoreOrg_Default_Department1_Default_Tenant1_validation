// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671695106646/Outdoor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Outdoor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OutdoorListener is a complete listener for a parse tree produced by OutdoorParser.
type OutdoorListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOutdoor is called when entering the outdoor production.
	EnterOutdoor(c *OutdoorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOutdoor is called when exiting the outdoor production.
	ExitOutdoor(c *OutdoorContext)
}
