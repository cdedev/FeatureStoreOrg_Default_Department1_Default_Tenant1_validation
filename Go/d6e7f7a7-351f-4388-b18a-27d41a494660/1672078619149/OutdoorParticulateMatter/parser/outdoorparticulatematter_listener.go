// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672078619149/OutdoorParticulateMatter.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OutdoorParticulateMatter

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OutdoorParticulateMatterListener is a complete listener for a parse tree produced by OutdoorParticulateMatterParser.
type OutdoorParticulateMatterListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOutdoorparticulatematter is called when entering the outdoorparticulatematter production.
	EnterOutdoorparticulatematter(c *OutdoorparticulatematterContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOutdoorparticulatematter is called when exiting the outdoorparticulatematter production.
	ExitOutdoorparticulatematter(c *OutdoorparticulatematterContext)
}
