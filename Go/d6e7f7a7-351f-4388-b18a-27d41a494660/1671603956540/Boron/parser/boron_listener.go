// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603956540/Boron.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Boron

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BoronListener is a complete listener for a parse tree produced by BoronParser.
type BoronListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBoron is called when entering the boron production.
	EnterBoron(c *BoronContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBoron is called when exiting the boron production.
	ExitBoron(c *BoronContext)
}
