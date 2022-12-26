// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077012015/MarMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MarMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MarMaxtempListener is a complete listener for a parse tree produced by MarMaxtempParser.
type MarMaxtempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMarmaxtemp is called when entering the marmaxtemp production.
	EnterMarmaxtemp(c *MarmaxtempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMarmaxtemp is called when exiting the marmaxtemp production.
	ExitMarmaxtemp(c *MarmaxtempContext)
}
