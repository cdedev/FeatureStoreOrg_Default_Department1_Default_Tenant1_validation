// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671598824557/Coal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Coal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CoalListener is a complete listener for a parse tree produced by CoalParser.
type CoalListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCoal is called when entering the coal production.
	EnterCoal(c *CoalContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCoal is called when exiting the coal production.
	ExitCoal(c *CoalContext)
}
