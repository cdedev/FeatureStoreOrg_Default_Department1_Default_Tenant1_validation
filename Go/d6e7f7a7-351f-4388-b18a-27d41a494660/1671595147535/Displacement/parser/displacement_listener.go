// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671595147535/Displacement.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Displacement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DisplacementListener is a complete listener for a parse tree produced by DisplacementParser.
type DisplacementListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDisplacement is called when entering the displacement production.
	EnterDisplacement(c *DisplacementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDisplacement is called when exiting the displacement production.
	ExitDisplacement(c *DisplacementContext)
}
