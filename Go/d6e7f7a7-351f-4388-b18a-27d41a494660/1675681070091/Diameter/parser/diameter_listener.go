// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675681070091/Diameter.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diameter

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiameterListener is a complete listener for a parse tree produced by DiameterParser.
type DiameterListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDiameter is called when entering the diameter production.
	EnterDiameter(c *DiameterContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDiameter is called when exiting the diameter production.
	ExitDiameter(c *DiameterContext)
}
