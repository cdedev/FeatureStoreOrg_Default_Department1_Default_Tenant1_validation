// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673856126889/Acceleration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Acceleration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AccelerationListener is a complete listener for a parse tree produced by AccelerationParser.
type AccelerationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAcceleration is called when entering the acceleration production.
	EnterAcceleration(c *AccelerationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAcceleration is called when exiting the acceleration production.
	ExitAcceleration(c *AccelerationContext)
}
