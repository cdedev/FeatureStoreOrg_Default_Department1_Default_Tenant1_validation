// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669983532769/EngineSize.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EngineSize

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EngineSizeListener is a complete listener for a parse tree produced by EngineSizeParser.
type EngineSizeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEnginesize is called when entering the enginesize production.
	EnterEnginesize(c *EnginesizeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEnginesize is called when exiting the enginesize production.
	ExitEnginesize(c *EnginesizeContext)
}
