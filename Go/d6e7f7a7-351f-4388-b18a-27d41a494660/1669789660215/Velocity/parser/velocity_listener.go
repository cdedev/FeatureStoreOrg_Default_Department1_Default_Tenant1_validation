// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669789660215/Velocity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Velocity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VelocityListener is a complete listener for a parse tree produced by VelocityParser.
type VelocityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVelocity is called when entering the velocity production.
	EnterVelocity(c *VelocityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVelocity is called when exiting the velocity production.
	ExitVelocity(c *VelocityContext)
}
