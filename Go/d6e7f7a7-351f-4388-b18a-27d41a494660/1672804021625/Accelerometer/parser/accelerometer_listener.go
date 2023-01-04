// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672804021625/Accelerometer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Accelerometer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AccelerometerListener is a complete listener for a parse tree produced by AccelerometerParser.
type AccelerometerListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAccelerometer is called when entering the accelerometer production.
	EnterAccelerometer(c *AccelerometerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAccelerometer is called when exiting the accelerometer production.
	ExitAccelerometer(c *AccelerometerContext)
}
