// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675743437296/Downloads.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Downloads

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DownloadsListener is a complete listener for a parse tree produced by DownloadsParser.
type DownloadsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDownloads is called when entering the downloads production.
	EnterDownloads(c *DownloadsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDownloads is called when exiting the downloads production.
	ExitDownloads(c *DownloadsContext)
}
