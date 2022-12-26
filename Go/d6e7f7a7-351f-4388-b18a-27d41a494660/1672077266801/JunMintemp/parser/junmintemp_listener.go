// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077266801/JunMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JunMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JunMintempListener is a complete listener for a parse tree produced by JunMintempParser.
type JunMintempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterJunmintemp is called when entering the junmintemp production.
	EnterJunmintemp(c *JunmintempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitJunmintemp is called when exiting the junmintemp production.
	ExitJunmintemp(c *JunmintempContext)
}
