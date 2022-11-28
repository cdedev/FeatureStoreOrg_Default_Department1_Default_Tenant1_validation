// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656772997/Pressure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pressure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PressureListener is a complete listener for a parse tree produced by PressureParser.
type PressureListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPressure is called when entering the pressure production.
	EnterPressure(c *PressureContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPressure is called when exiting the pressure production.
	ExitPressure(c *PressureContext)
}
