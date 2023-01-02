// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637624734/HIndex.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // HIndex

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HIndexListener is a complete listener for a parse tree produced by HIndexParser.
type HIndexListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHindex is called when entering the hindex production.
	EnterHindex(c *HindexContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHindex is called when exiting the hindex production.
	ExitHindex(c *HindexContext)
}
