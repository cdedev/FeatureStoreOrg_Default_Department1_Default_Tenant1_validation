// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671517657930/RadiativePower.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RadiativePower

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RadiativePowerListener is a complete listener for a parse tree produced by RadiativePowerParser.
type RadiativePowerListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRadiativepower is called when entering the radiativepower production.
	EnterRadiativepower(c *RadiativepowerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRadiativepower is called when exiting the radiativepower production.
	ExitRadiativepower(c *RadiativepowerContext)
}
