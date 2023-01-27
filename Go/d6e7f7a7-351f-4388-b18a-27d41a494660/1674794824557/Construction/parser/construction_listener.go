// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674794824557/Construction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Construction

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConstructionListener is a complete listener for a parse tree produced by ConstructionParser.
type ConstructionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConstruction is called when entering the construction production.
	EnterConstruction(c *ConstructionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConstruction is called when exiting the construction production.
	ExitConstruction(c *ConstructionContext)
}
