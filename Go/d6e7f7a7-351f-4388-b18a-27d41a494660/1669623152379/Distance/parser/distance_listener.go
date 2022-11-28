// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623152379/Distance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Distance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DistanceListener is a complete listener for a parse tree produced by DistanceParser.
type DistanceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDistance is called when entering the distance production.
	EnterDistance(c *DistanceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDistance is called when exiting the distance production.
	ExitDistance(c *DistanceContext)
}
