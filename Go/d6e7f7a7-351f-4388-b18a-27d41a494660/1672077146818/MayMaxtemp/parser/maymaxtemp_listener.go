// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077146818/MayMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MayMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MayMaxtempListener is a complete listener for a parse tree produced by MayMaxtempParser.
type MayMaxtempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMaymaxtemp is called when entering the maymaxtemp production.
	EnterMaymaxtemp(c *MaymaxtempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMaymaxtemp is called when exiting the maymaxtemp production.
	ExitMaymaxtemp(c *MaymaxtempContext)
}
