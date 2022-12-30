// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377318029/DualSim.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DualSim

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DualSimListener is a complete listener for a parse tree produced by DualSimParser.
type DualSimListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDualsim is called when entering the dualsim production.
	EnterDualsim(c *DualsimContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDualsim is called when exiting the dualsim production.
	ExitDualsim(c *DualsimContext)
}
