// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671203166456/Terrace.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Terrace

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TerraceListener is a complete listener for a parse tree produced by TerraceParser.
type TerraceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTerrace is called when entering the terrace production.
	EnterTerrace(c *TerraceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTerrace is called when exiting the terrace production.
	ExitTerrace(c *TerraceContext)
}
