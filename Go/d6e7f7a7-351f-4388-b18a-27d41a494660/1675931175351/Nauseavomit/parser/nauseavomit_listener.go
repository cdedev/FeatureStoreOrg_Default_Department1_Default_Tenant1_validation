// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675931175351/Nauseavomit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nauseavomit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NauseavomitListener is a complete listener for a parse tree produced by NauseavomitParser.
type NauseavomitListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNauseavomit is called when entering the nauseavomit production.
	EnterNauseavomit(c *NauseavomitContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNauseavomit is called when exiting the nauseavomit production.
	ExitNauseavomit(c *NauseavomitContext)
}
