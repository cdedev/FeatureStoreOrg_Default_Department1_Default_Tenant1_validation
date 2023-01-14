// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673676368835/Radius.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Radius

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RadiusListener is a complete listener for a parse tree produced by RadiusParser.
type RadiusListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRadius is called when entering the radius production.
	EnterRadius(c *RadiusContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRadius is called when exiting the radius production.
	ExitRadius(c *RadiusContext)
}
