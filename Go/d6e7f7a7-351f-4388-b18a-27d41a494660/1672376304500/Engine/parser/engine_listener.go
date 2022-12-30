// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376304500/Engine.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Engine

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EngineListener is a complete listener for a parse tree produced by EngineParser.
type EngineListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEngine is called when entering the engine production.
	EnterEngine(c *EngineContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEngine is called when exiting the engine production.
	ExitEngine(c *EngineContext)
}
