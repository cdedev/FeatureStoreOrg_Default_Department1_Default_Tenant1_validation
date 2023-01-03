// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672718641270/Key.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Key

import "github.com/antlr/antlr4/runtime/Go/antlr"

// KeyListener is a complete listener for a parse tree produced by KeyParser.
type KeyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)
}
