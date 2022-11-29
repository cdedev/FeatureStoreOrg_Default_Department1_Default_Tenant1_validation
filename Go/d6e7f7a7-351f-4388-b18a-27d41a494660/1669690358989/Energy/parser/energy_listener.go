// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690358989/Energy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Energy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EnergyListener is a complete listener for a parse tree produced by EnergyParser.
type EnergyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEnergy is called when entering the energy production.
	EnterEnergy(c *EnergyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEnergy is called when exiting the energy production.
	ExitEnergy(c *EnergyContext)
}
