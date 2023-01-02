// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672635660795/GameId.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GameId

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GameIdListener is a complete listener for a parse tree produced by GameIdParser.
type GameIdListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGameid is called when entering the gameid production.
	EnterGameid(c *GameidContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGameid is called when exiting the gameid production.
	ExitGameid(c *GameidContext)
}
