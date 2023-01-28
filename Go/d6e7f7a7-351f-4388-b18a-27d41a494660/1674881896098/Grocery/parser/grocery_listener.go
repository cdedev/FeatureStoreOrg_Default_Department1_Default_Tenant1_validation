// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674881896098/Grocery.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Grocery

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GroceryListener is a complete listener for a parse tree produced by GroceryParser.
type GroceryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGrocery is called when entering the grocery production.
	EnterGrocery(c *GroceryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGrocery is called when exiting the grocery production.
	ExitGrocery(c *GroceryContext)
}
