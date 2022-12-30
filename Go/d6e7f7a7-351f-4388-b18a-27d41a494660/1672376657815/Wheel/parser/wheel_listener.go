// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376657815/Wheel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wheel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WheelListener is a complete listener for a parse tree produced by WheelParser.
type WheelListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWheel is called when entering the wheel production.
	EnterWheel(c *WheelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWheel is called when exiting the wheel production.
	ExitWheel(c *WheelContext)
}
