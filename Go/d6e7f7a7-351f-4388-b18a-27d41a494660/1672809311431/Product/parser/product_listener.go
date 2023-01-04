// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672809311431/Product.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Product

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ProductListener is a complete listener for a parse tree produced by ProductParser.
type ProductListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterProduct is called when entering the product production.
	EnterProduct(c *ProductContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitProduct is called when exiting the product production.
	ExitProduct(c *ProductContext)
}
