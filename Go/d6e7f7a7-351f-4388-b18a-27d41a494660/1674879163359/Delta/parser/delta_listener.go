// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674879163359/Delta.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Delta

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DeltaListener is a complete listener for a parse tree produced by DeltaParser.
type DeltaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDelta is called when entering the delta production.
	EnterDelta(c *DeltaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDelta is called when exiting the delta production.
	ExitDelta(c *DeltaContext)
}
