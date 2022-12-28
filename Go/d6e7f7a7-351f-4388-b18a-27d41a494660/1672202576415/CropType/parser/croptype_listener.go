// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202576415/CropType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CropType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CropTypeListener is a complete listener for a parse tree produced by CropTypeParser.
type CropTypeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCroptype is called when entering the croptype production.
	EnterCroptype(c *CroptypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCroptype is called when exiting the croptype production.
	ExitCroptype(c *CroptypeContext)
}
