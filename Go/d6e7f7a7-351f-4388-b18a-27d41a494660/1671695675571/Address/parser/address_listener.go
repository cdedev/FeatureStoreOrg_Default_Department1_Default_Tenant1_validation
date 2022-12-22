// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671695675571/Address.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Address

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AddressListener is a complete listener for a parse tree produced by AddressParser.
type AddressListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAddress is called when entering the address production.
	EnterAddress(c *AddressContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAddress is called when exiting the address production.
	ExitAddress(c *AddressContext)
}
