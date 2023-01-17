// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925719765/Milk.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Milk

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MilkListener is a complete listener for a parse tree produced by MilkParser.
type MilkListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMilk is called when entering the milk production.
	EnterMilk(c *MilkContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMilk is called when exiting the milk production.
	ExitMilk(c *MilkContext)
}
