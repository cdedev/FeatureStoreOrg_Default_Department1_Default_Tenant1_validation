// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675078116689/Shape.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Shape

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ShapeListener is a complete listener for a parse tree produced by ShapeParser.
type ShapeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterShape is called when entering the shape production.
	EnterShape(c *ShapeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitShape is called when exiting the shape production.
	ExitShape(c *ShapeContext)
}
