// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675703060989/Points.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Points

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PointsListener is a complete listener for a parse tree produced by PointsParser.
type PointsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPoints is called when entering the points production.
	EnterPoints(c *PointsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPoints is called when exiting the points production.
	ExitPoints(c *PointsContext)
}
