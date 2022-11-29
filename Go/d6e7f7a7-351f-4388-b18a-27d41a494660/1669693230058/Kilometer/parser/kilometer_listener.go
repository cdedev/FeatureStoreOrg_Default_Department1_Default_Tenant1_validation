// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669693230058/Kilometer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Kilometer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// KilometerListener is a complete listener for a parse tree produced by KilometerParser.
type KilometerListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterKilometer is called when entering the kilometer production.
	EnterKilometer(c *KilometerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitKilometer is called when exiting the kilometer production.
	ExitKilometer(c *KilometerContext)
}
