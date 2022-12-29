// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672288200811/Desibel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Desibel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DesibelListener is a complete listener for a parse tree produced by DesibelParser.
type DesibelListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDesibel is called when entering the desibel production.
	EnterDesibel(c *DesibelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDesibel is called when exiting the desibel production.
	ExitDesibel(c *DesibelContext)
}
