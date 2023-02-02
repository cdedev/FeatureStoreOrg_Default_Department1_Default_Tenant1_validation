// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675306627938/Fat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FatListener is a complete listener for a parse tree produced by FatParser.
type FatListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFat is called when entering the fat production.
	EnterFat(c *FatContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFat is called when exiting the fat production.
	ExitFat(c *FatContext)
}
