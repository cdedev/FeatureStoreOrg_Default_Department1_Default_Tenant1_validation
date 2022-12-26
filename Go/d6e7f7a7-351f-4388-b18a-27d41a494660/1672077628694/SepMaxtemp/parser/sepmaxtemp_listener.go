// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077628694/SepMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SepMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SepMaxtempListener is a complete listener for a parse tree produced by SepMaxtempParser.
type SepMaxtempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSepmaxtemp is called when entering the sepmaxtemp production.
	EnterSepmaxtemp(c *SepmaxtempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSepmaxtemp is called when exiting the sepmaxtemp production.
	ExitSepmaxtemp(c *SepmaxtempContext)
}
