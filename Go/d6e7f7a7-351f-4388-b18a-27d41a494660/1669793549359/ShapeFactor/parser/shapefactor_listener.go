// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669793549359/ShapeFactor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ShapeFactor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ShapeFactorListener is a complete listener for a parse tree produced by ShapeFactorParser.
type ShapeFactorListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterShapefactor is called when entering the shapefactor production.
	EnterShapefactor(c *ShapefactorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitShapefactor is called when exiting the shapefactor production.
	ExitShapefactor(c *ShapefactorContext)
}
