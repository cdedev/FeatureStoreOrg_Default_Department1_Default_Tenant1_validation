// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637572205/Documents.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Documents

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DocumentsListener is a complete listener for a parse tree produced by DocumentsParser.
type DocumentsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDocuments is called when entering the documents production.
	EnterDocuments(c *DocumentsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDocuments is called when exiting the documents production.
	ExitDocuments(c *DocumentsContext)
}
