// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676609525407/Quality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Quality

import "github.com/antlr/antlr4/runtime/Go/antlr"

// QualityListener is a complete listener for a parse tree produced by QualityParser.
type QualityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterQuality is called when entering the quality production.
	EnterQuality(c *QualityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitQuality is called when exiting the quality production.
	ExitQuality(c *QualityContext)
}
