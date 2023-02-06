// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675680683851/Material.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Material

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MaterialListener is a complete listener for a parse tree produced by MaterialParser.
type MaterialListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMaterial is called when entering the material production.
	EnterMaterial(c *MaterialContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMaterial is called when exiting the material production.
	ExitMaterial(c *MaterialContext)
}
