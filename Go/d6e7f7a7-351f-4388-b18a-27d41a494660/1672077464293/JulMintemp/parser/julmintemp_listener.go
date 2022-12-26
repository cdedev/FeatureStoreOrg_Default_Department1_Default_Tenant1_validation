// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077464293/JulMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JulMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JulMintempListener is a complete listener for a parse tree produced by JulMintempParser.
type JulMintempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterJulmintemp is called when entering the julmintemp production.
	EnterJulmintemp(c *JulmintempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitJulmintemp is called when exiting the julmintemp production.
	ExitJulmintemp(c *JulmintempContext)
}
