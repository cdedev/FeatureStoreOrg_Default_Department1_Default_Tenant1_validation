// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671208580886/Surface.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Surface

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SurfaceListener is a complete listener for a parse tree produced by SurfaceParser.
type SurfaceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSurface is called when entering the surface production.
	EnterSurface(c *SurfaceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSurface is called when exiting the surface production.
	ExitSurface(c *SurfaceContext)
}
