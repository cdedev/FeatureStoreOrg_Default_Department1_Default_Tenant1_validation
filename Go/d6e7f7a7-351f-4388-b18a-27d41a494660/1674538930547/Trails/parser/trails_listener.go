// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674538930547/Trails.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Trails

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TrailsListener is a complete listener for a parse tree produced by TrailsParser.
type TrailsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTrails is called when entering the trails production.
	EnterTrails(c *TrailsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTrails is called when exiting the trails production.
	ExitTrails(c *TrailsContext)
}
