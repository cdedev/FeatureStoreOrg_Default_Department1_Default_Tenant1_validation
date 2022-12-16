// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671186145445/WaterLevel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WaterLevel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WaterLevelListener is a complete listener for a parse tree produced by WaterLevelParser.
type WaterLevelListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWaterlevel is called when entering the waterlevel production.
	EnterWaterlevel(c *WaterlevelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWaterlevel is called when exiting the waterlevel production.
	ExitWaterlevel(c *WaterlevelContext)
}
