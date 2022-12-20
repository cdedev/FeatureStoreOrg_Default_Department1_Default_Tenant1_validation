// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671523287895/Season.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Season

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SeasonListener is a complete listener for a parse tree produced by SeasonParser.
type SeasonListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSeason is called when entering the season production.
	EnterSeason(c *SeasonContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSeason is called when exiting the season production.
	ExitSeason(c *SeasonContext)
}
