// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669691961999/Thickness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Thickness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ThicknessListener is a complete listener for a parse tree produced by ThicknessParser.
type ThicknessListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterThickness is called when entering the thickness production.
	EnterThickness(c *ThicknessContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitThickness is called when exiting the thickness production.
	ExitThickness(c *ThicknessContext)
}
