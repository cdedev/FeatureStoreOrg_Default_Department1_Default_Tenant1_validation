// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377702280/FourG.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FourG

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FourGListener is a complete listener for a parse tree produced by FourGParser.
type FourGListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFourg is called when entering the fourg production.
	EnterFourg(c *FourgContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFourg is called when exiting the fourg production.
	ExitFourg(c *FourgContext)
}
