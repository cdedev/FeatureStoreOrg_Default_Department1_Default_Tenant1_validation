// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077708967/SepMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SepMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SepMintempListener is a complete listener for a parse tree produced by SepMintempParser.
type SepMintempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSepmintemp is called when entering the sepmintemp production.
	EnterSepmintemp(c *SepmintempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSepmintemp is called when exiting the sepmintemp production.
	ExitSepmintemp(c *SepmintempContext)
}
