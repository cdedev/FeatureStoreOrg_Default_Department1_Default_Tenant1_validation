// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675758417557/AvgDailyPageviews.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AvgDailyPageviews

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AvgDailyPageviewsListener is a complete listener for a parse tree produced by AvgDailyPageviewsParser.
type AvgDailyPageviewsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAvgdailypageviews is called when entering the avgdailypageviews production.
	EnterAvgdailypageviews(c *AvgdailypageviewsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAvgdailypageviews is called when exiting the avgdailypageviews production.
	ExitAvgdailypageviews(c *AvgdailypageviewsContext)
}
