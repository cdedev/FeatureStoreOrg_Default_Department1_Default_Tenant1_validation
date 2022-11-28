// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669655797520/Heartrate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Heartrate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HeartrateListener is a complete listener for a parse tree produced by HeartrateParser.
type HeartrateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHeartrate is called when entering the heartrate production.
	EnterHeartrate(c *HeartrateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHeartrate is called when exiting the heartrate production.
	ExitHeartrate(c *HeartrateContext)
}
