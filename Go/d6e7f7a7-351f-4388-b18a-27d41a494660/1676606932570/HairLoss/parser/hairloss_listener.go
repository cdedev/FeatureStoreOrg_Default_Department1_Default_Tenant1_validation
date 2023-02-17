// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676606932570/HairLoss.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // HairLoss

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HairLossListener is a complete listener for a parse tree produced by HairLossParser.
type HairLossListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHairloss is called when entering the hairloss production.
	EnterHairloss(c *HairlossContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHairloss is called when exiting the hairloss production.
	ExitHairloss(c *HairlossContext)
}
