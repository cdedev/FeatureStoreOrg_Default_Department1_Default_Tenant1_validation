// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672462254377/Elevation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Elevation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ElevationListener is a complete listener for a parse tree produced by ElevationParser.
type ElevationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterElevation is called when entering the elevation production.
	EnterElevation(c *ElevationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitElevation is called when exiting the elevation production.
	ExitElevation(c *ElevationContext)
}
