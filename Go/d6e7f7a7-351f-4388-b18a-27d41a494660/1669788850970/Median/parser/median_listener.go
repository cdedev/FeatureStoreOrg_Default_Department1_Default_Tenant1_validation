// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669788850970/Median.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Median

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MedianListener is a complete listener for a parse tree produced by MedianParser.
type MedianListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMedian is called when entering the median production.
	EnterMedian(c *MedianContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMedian is called when exiting the median production.
	ExitMedian(c *MedianContext)
}
