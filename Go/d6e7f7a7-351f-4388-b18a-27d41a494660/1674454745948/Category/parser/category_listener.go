// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674454745948/Category.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Category

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CategoryListener is a complete listener for a parse tree produced by CategoryParser.
type CategoryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCategory is called when entering the category production.
	EnterCategory(c *CategoryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCategory is called when exiting the category production.
	ExitCategory(c *CategoryContext)
}
