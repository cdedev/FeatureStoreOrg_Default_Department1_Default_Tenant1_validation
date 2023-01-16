// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878091669/Dysarthria.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dysarthria

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DysarthriaListener is a complete listener for a parse tree produced by DysarthriaParser.
type DysarthriaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDysarthria is called when entering the dysarthria production.
	EnterDysarthria(c *DysarthriaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDysarthria is called when exiting the dysarthria production.
	ExitDysarthria(c *DysarthriaContext)
}
