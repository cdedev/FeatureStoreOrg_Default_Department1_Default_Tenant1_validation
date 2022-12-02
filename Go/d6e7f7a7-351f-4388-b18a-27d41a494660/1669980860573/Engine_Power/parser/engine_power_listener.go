// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669980860573/Engine_Power.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Engine_Power

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Engine_PowerListener is a complete listener for a parse tree produced by Engine_PowerParser.
type Engine_PowerListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEngine_power is called when entering the engine_power production.
	EnterEngine_power(c *Engine_powerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEngine_power is called when exiting the engine_power production.
	ExitEngine_power(c *Engine_powerContext)
}
