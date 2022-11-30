// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669789074227/Gas.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gas

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GasListener is a complete listener for a parse tree produced by GasParser.
type GasListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGas is called when entering the gas production.
	EnterGas(c *GasContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGas is called when exiting the gas production.
	ExitGas(c *GasContext)
}
