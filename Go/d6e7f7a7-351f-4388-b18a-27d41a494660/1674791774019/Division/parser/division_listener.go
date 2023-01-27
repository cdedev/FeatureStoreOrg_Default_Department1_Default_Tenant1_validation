// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674791774019/Division.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Division

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DivisionListener is a complete listener for a parse tree produced by DivisionParser.
type DivisionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDivision is called when entering the division production.
	EnterDivision(c *DivisionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDivision is called when exiting the division production.
	ExitDivision(c *DivisionContext)
}
