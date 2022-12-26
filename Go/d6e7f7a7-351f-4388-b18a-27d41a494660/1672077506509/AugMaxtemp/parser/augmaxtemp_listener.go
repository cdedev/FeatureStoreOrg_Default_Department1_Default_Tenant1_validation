// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077506509/AugMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AugMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AugMaxtempListener is a complete listener for a parse tree produced by AugMaxtempParser.
type AugMaxtempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAugmaxtemp is called when entering the augmaxtemp production.
	EnterAugmaxtemp(c *AugmaxtempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAugmaxtemp is called when exiting the augmaxtemp production.
	ExitAugmaxtemp(c *AugmaxtempContext)
}
