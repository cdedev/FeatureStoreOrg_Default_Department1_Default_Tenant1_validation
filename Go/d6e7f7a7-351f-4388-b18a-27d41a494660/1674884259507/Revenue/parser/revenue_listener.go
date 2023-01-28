// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674884259507/Revenue.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Revenue

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RevenueListener is a complete listener for a parse tree produced by RevenueParser.
type RevenueListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRevenue is called when entering the revenue production.
	EnterRevenue(c *RevenueContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRevenue is called when exiting the revenue production.
	ExitRevenue(c *RevenueContext)
}
