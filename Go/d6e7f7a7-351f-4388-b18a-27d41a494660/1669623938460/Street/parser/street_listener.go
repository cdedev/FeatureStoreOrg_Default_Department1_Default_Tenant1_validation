// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623938460/Street.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Street

import "github.com/antlr/antlr4/runtime/Go/antlr"

// StreetListener is a complete listener for a parse tree produced by StreetParser.
type StreetListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterStreet is called when entering the street production.
	EnterStreet(c *StreetContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitStreet is called when exiting the street production.
	ExitStreet(c *StreetContext)
}
