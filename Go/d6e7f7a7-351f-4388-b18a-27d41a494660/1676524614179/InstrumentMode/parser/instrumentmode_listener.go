// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524614179/InstrumentMode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // InstrumentMode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InstrumentModeListener is a complete listener for a parse tree produced by InstrumentModeParser.
type InstrumentModeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInstrumentmode is called when entering the instrumentmode production.
	EnterInstrumentmode(c *InstrumentmodeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInstrumentmode is called when exiting the instrumentmode production.
	ExitInstrumentmode(c *InstrumentmodeContext)
}
