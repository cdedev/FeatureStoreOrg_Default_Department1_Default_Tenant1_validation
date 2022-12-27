// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672112976625/DriverPresent.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DriverPresent

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DriverPresentListener is a complete listener for a parse tree produced by DriverPresentParser.
type DriverPresentListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDriverpresent is called when entering the driverpresent production.
	EnterDriverpresent(c *DriverpresentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDriverpresent is called when exiting the driverpresent production.
	ExitDriverpresent(c *DriverpresentContext)
}
