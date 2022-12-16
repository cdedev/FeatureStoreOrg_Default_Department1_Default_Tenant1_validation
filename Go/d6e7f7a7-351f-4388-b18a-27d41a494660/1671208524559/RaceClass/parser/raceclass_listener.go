// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671208524559/RaceClass.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RaceClass

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RaceClassListener is a complete listener for a parse tree produced by RaceClassParser.
type RaceClassListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRaceclass is called when entering the raceclass production.
	EnterRaceclass(c *RaceclassContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRaceclass is called when exiting the raceclass production.
	ExitRaceclass(c *RaceclassContext)
}
