// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673667498609/Domestic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Domestic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DomesticListener is a complete listener for a parse tree produced by DomesticParser.
type DomesticListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDomestic is called when entering the domestic production.
	EnterDomestic(c *DomesticContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDomestic is called when exiting the domestic production.
	ExitDomestic(c *DomesticContext)
}
