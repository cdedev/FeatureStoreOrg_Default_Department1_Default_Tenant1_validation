// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674666977432/Race.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Race

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RaceListener is a complete listener for a parse tree produced by RaceParser.
type RaceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRace is called when entering the race production.
	EnterRace(c *RaceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRace is called when exiting the race production.
	ExitRace(c *RaceContext)
}
