// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669699346197/Perimeter.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Perimeter

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PerimeterListener is a complete listener for a parse tree produced by PerimeterParser.
type PerimeterListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPerimeter is called when entering the perimeter production.
	EnterPerimeter(c *PerimeterContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPerimeter is called when exiting the perimeter production.
	ExitPerimeter(c *PerimeterContext)
}
