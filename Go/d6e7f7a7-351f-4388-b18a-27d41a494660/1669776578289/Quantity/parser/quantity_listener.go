// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669776578289/Quantity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Quantity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// QuantityListener is a complete listener for a parse tree produced by QuantityParser.
type QuantityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterQuantity is called when entering the quantity production.
	EnterQuantity(c *QuantityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitQuantity is called when exiting the quantity production.
	ExitQuantity(c *QuantityContext)
}
