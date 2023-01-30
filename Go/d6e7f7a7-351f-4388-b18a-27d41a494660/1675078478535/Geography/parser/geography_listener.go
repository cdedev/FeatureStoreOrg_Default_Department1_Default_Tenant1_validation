// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675078478535/Geography.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Geography

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GeographyListener is a complete listener for a parse tree produced by GeographyParser.
type GeographyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGeography is called when entering the geography production.
	EnterGeography(c *GeographyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGeography is called when exiting the geography production.
	ExitGeography(c *GeographyContext)
}
