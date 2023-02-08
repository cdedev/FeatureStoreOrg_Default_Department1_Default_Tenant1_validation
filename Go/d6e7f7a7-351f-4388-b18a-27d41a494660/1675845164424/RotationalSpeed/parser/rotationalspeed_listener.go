// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675845164424/RotationalSpeed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RotationalSpeed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RotationalSpeedListener is a complete listener for a parse tree produced by RotationalSpeedParser.
type RotationalSpeedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRotationalspeed is called when entering the rotationalspeed production.
	EnterRotationalspeed(c *RotationalspeedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRotationalspeed is called when exiting the rotationalspeed production.
	ExitRotationalspeed(c *RotationalspeedContext)
}
