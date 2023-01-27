// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674793042480/Wards.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wards

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WardsListener is a complete listener for a parse tree produced by WardsParser.
type WardsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWards is called when entering the wards production.
	EnterWards(c *WardsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWards is called when exiting the wards production.
	ExitWards(c *WardsContext)
}
