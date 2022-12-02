// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669983780063/CO2Emissions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CO2Emissions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CO2EmissionsListener is a complete listener for a parse tree produced by CO2EmissionsParser.
type CO2EmissionsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCo2emissions is called when entering the co2emissions production.
	EnterCo2emissions(c *Co2emissionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCo2emissions is called when exiting the co2emissions production.
	ExitCo2emissions(c *Co2emissionsContext)
}
