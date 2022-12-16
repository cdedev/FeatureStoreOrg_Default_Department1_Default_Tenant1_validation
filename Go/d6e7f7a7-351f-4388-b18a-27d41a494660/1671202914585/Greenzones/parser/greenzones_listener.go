// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671202914585/Greenzones.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Greenzones

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GreenzonesListener is a complete listener for a parse tree produced by GreenzonesParser.
type GreenzonesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGreenzones is called when entering the greenzones production.
	EnterGreenzones(c *GreenzonesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGreenzones is called when exiting the greenzones production.
	ExitGreenzones(c *GreenzonesContext)
}
