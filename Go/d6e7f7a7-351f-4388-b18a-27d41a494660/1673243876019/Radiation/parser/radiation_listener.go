// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673243876019/Radiation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Radiation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RadiationListener is a complete listener for a parse tree produced by RadiationParser.
type RadiationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRadiation is called when entering the radiation production.
	EnterRadiation(c *RadiationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRadiation is called when exiting the radiation production.
	ExitRadiation(c *RadiationContext)
}
