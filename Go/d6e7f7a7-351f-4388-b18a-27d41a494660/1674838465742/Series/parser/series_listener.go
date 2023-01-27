// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674838465742/Series.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Series

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SeriesListener is a complete listener for a parse tree produced by SeriesParser.
type SeriesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSeries is called when entering the series production.
	EnterSeries(c *SeriesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSeries is called when exiting the series production.
	ExitSeries(c *SeriesContext)
}
