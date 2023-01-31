// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675145305020/Bytes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bytes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BytesListener is a complete listener for a parse tree produced by BytesParser.
type BytesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBytes is called when entering the bytes production.
	EnterBytes(c *BytesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBytes is called when exiting the bytes production.
	ExitBytes(c *BytesContext)
}
