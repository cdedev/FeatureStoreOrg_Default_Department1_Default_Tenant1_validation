// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671767493572/Voltage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Voltage

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VoltageListener is a complete listener for a parse tree produced by VoltageParser.
type VoltageListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVoltage is called when entering the voltage production.
	EnterVoltage(c *VoltageContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVoltage is called when exiting the voltage production.
	ExitVoltage(c *VoltageContext)
}
