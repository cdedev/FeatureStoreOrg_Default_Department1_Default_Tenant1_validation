// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077553967/AugMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AugMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AugMintempListener is a complete listener for a parse tree produced by AugMintempParser.
type AugMintempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAugmintemp is called when entering the augmintemp production.
	EnterAugmintemp(c *AugmintempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAugmintemp is called when exiting the augmintemp production.
	ExitAugmintemp(c *AugmintempContext)
}
