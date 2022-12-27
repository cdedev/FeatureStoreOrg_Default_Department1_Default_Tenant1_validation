// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672112566753/Games.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Games

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GamesListener is a complete listener for a parse tree produced by GamesParser.
type GamesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGames is called when entering the games production.
	EnterGames(c *GamesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGames is called when exiting the games production.
	ExitGames(c *GamesContext)
}
