// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675241833178/DateOfDeath.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DateOfDeath

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DateOfDeathListener is a complete listener for a parse tree produced by DateOfDeathParser.
type DateOfDeathListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDateofdeath is called when entering the dateofdeath production.
	EnterDateofdeath(c *DateofdeathContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDateofdeath is called when exiting the dateofdeath production.
	ExitDateofdeath(c *DateofdeathContext)
}
