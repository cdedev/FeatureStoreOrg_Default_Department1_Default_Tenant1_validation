// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284768398/CartoonReference.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CartoonReference

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CartoonReferenceListener is a complete listener for a parse tree produced by CartoonReferenceParser.
type CartoonReferenceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCartoonreference is called when entering the cartoonreference production.
	EnterCartoonreference(c *CartoonreferenceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCartoonreference is called when exiting the cartoonreference production.
	ExitCartoonreference(c *CartoonreferenceContext)
}
