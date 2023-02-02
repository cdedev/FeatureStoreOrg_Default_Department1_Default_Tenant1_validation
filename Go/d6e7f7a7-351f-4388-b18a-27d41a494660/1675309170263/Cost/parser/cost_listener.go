// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675309170263/Cost.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cost

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CostListener is a complete listener for a parse tree produced by CostParser.
type CostListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCost is called when entering the cost production.
	EnterCost(c *CostContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCost is called when exiting the cost production.
	ExitCost(c *CostContext)
}
