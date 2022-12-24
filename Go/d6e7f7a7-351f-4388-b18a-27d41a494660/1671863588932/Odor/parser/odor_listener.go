// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671863588932/Odor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Odor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OdorListener is a complete listener for a parse tree produced by OdorParser.
type OdorListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOdor is called when entering the odor production.
	EnterOdor(c *OdorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOdor is called when exiting the odor production.
	ExitOdor(c *OdorContext)
}
