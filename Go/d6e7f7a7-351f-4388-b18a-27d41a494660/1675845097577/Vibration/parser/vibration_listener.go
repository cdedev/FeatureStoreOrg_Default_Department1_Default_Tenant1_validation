// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675845097577/Vibration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vibration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VibrationListener is a complete listener for a parse tree produced by VibrationParser.
type VibrationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVibration is called when entering the vibration production.
	EnterVibration(c *VibrationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVibration is called when exiting the vibration production.
	ExitVibration(c *VibrationContext)
}
