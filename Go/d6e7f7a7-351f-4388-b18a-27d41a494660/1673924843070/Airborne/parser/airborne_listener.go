// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673924843070/Airborne.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Airborne

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AirborneListener is a complete listener for a parse tree produced by AirborneParser.
type AirborneListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAirborne is called when entering the airborne production.
	EnterAirborne(c *AirborneContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAirborne is called when exiting the airborne production.
	ExitAirborne(c *AirborneContext)
}
