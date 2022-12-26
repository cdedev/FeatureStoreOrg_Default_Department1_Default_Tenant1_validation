// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672073289495/Level.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Level

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LevelListener is a complete listener for a parse tree produced by LevelParser.
type LevelListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLevel is called when entering the level production.
	EnterLevel(c *LevelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLevel is called when exiting the level production.
	ExitLevel(c *LevelContext)
}
