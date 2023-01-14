// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673676416291/Smoothness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Smoothness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SmoothnessListener is a complete listener for a parse tree produced by SmoothnessParser.
type SmoothnessListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSmoothness is called when entering the smoothness production.
	EnterSmoothness(c *SmoothnessContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSmoothness is called when exiting the smoothness production.
	ExitSmoothness(c *SmoothnessContext)
}
