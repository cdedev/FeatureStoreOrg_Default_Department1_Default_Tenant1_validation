// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235164422/Openness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Openness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OpennessListener is a complete listener for a parse tree produced by OpennessParser.
type OpennessListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOpenness is called when entering the openness production.
	EnterOpenness(c *OpennessContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOpenness is called when exiting the openness production.
	ExitOpenness(c *OpennessContext)
}
