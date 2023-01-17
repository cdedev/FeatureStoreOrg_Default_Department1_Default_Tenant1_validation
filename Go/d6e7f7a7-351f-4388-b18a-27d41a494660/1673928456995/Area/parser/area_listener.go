// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673928456995/Area.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Area

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AreaListener is a complete listener for a parse tree produced by AreaParser.
type AreaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterArea is called when entering the area production.
	EnterArea(c *AreaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitArea is called when exiting the area production.
	ExitArea(c *AreaContext)
}
