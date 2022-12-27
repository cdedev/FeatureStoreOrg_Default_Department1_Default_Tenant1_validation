// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672115062165/Traffic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Traffic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TrafficListener is a complete listener for a parse tree produced by TrafficParser.
type TrafficListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTraffic is called when entering the traffic production.
	EnterTraffic(c *TrafficContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTraffic is called when exiting the traffic production.
	ExitTraffic(c *TrafficContext)
}
