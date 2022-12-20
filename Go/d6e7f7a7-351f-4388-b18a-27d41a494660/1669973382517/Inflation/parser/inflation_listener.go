// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669973382517/Inflation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Inflation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InflationListener is a complete listener for a parse tree produced by InflationParser.
type InflationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInflation is called when entering the inflation production.
	EnterInflation(c *InflationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInflation is called when exiting the inflation production.
	ExitInflation(c *InflationContext)
}