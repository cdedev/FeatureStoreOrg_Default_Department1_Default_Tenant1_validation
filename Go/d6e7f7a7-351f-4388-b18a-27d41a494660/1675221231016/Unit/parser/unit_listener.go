// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675221231016/Unit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Unit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// UnitListener is a complete listener for a parse tree produced by UnitParser.
type UnitListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)
}
