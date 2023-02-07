// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675758509632/Impressions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Impressions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ImpressionsListener is a complete listener for a parse tree produced by ImpressionsParser.
type ImpressionsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterImpressions is called when entering the impressions production.
	EnterImpressions(c *ImpressionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitImpressions is called when exiting the impressions production.
	ExitImpressions(c *ImpressionsContext)
}
