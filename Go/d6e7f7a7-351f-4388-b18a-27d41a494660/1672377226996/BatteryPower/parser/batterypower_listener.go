// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377226996/BatteryPower.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BatteryPower

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BatteryPowerListener is a complete listener for a parse tree produced by BatteryPowerParser.
type BatteryPowerListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBatterypower is called when entering the batterypower production.
	EnterBatterypower(c *BatterypowerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBatterypower is called when exiting the batterypower production.
	ExitBatterypower(c *BatterypowerContext)
}
