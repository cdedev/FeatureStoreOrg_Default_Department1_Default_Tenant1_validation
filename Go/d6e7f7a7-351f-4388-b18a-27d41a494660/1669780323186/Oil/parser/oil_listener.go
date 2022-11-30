// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669780323186/Oil.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Oil

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OilListener is a complete listener for a parse tree produced by OilParser.
type OilListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOil is called when entering the oil production.
	EnterOil(c *OilContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOil is called when exiting the oil production.
	ExitOil(c *OilContext)
}
