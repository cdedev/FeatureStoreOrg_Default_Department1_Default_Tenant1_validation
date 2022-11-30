// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669797355387/Torque.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Torque

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TorqueListener is a complete listener for a parse tree produced by TorqueParser.
type TorqueListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTorque is called when entering the torque production.
	EnterTorque(c *TorqueContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTorque is called when exiting the torque production.
	ExitTorque(c *TorqueContext)
}
