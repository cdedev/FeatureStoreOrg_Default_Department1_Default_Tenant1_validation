// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076977823/JanMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JanMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JanMintempListener is a complete listener for a parse tree produced by JanMintempParser.
type JanMintempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterJanmintemp is called when entering the janmintemp production.
	EnterJanmintemp(c *JanmintempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitJanmintemp is called when exiting the janmintemp production.
	ExitJanmintemp(c *JanmintempContext)
}
