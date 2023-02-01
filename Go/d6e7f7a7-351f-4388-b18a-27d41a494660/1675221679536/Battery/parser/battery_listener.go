// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675221679536/Battery.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Battery

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BatteryListener is a complete listener for a parse tree produced by BatteryParser.
type BatteryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBattery is called when entering the battery production.
	EnterBattery(c *BatteryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBattery is called when exiting the battery production.
	ExitBattery(c *BatteryContext)
}
