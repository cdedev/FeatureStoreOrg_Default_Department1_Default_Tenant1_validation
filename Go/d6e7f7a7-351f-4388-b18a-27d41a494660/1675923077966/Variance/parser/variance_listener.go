// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675923077966/Variance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Variance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VarianceListener is a complete listener for a parse tree produced by VarianceParser.
type VarianceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVariance is called when entering the variance production.
	EnterVariance(c *VarianceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVariance is called when exiting the variance production.
	ExitVariance(c *VarianceContext)
}
