// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675673022752/Correlation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Correlation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CorrelationListener is a complete listener for a parse tree produced by CorrelationParser.
type CorrelationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCorrelation is called when entering the correlation production.
	EnterCorrelation(c *CorrelationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCorrelation is called when exiting the correlation production.
	ExitCorrelation(c *CorrelationContext)
}
