// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077111598/AprMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AprMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AprMintempListener is a complete listener for a parse tree produced by AprMintempParser.
type AprMintempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAprmintemp is called when entering the aprmintemp production.
	EnterAprmintemp(c *AprmintempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAprmintemp is called when exiting the aprmintemp production.
	ExitAprmintemp(c *AprmintempContext)
}
