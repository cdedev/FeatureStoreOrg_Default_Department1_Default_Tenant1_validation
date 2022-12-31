// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672460775520/Structure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Structure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// StructureListener is a complete listener for a parse tree produced by StructureParser.
type StructureListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterStructure is called when entering the structure production.
	EnterStructure(c *StructureContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitStructure is called when exiting the structure production.
	ExitStructure(c *StructureContext)
}
