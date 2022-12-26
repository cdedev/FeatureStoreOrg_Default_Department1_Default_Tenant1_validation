// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077928289/NovMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NovMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NovMaxtempListener is a complete listener for a parse tree produced by NovMaxtempParser.
type NovMaxtempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNovmaxtemp is called when entering the novmaxtemp production.
	EnterNovmaxtemp(c *NovmaxtempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNovmaxtemp is called when exiting the novmaxtemp production.
	ExitNovmaxtemp(c *NovmaxtempContext)
}
