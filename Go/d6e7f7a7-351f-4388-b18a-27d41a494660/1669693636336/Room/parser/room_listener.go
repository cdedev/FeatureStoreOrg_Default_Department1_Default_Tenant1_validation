// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669693636336/Room.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Room

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RoomListener is a complete listener for a parse tree produced by RoomParser.
type RoomListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRoom is called when entering the room production.
	EnterRoom(c *RoomContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRoom is called when exiting the room production.
	ExitRoom(c *RoomContext)
}
