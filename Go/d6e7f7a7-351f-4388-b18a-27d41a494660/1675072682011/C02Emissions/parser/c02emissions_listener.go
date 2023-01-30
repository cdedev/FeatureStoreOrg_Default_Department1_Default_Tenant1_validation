// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072682011/C02Emissions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // C02Emissions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// C02EmissionsListener is a complete listener for a parse tree produced by C02EmissionsParser.
type C02EmissionsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterC02emissions is called when entering the c02emissions production.
	EnterC02emissions(c *C02emissionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitC02emissions is called when exiting the c02emissions production.
	ExitC02emissions(c *C02emissionsContext)
}
