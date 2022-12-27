// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119011639/Drops.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Drops

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DropsListener is a complete listener for a parse tree produced by DropsParser.
type DropsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDrops is called when entering the drops production.
	EnterDrops(c *DropsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDrops is called when exiting the drops production.
	ExitDrops(c *DropsContext)
}
