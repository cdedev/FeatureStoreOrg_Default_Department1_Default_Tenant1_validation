// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235590013/Axons.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Axons

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AxonsListener is a complete listener for a parse tree produced by AxonsParser.
type AxonsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAxons is called when entering the axons production.
	EnterAxons(c *AxonsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAxons is called when exiting the axons production.
	ExitAxons(c *AxonsContext)
}
