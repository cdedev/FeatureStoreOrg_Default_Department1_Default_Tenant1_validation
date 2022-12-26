// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077047300/MarMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MarMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MarMintempListener is a complete listener for a parse tree produced by MarMintempParser.
type MarMintempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMarmintemp is called when entering the marmintemp production.
	EnterMarmintemp(c *MarmintempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMarmintemp is called when exiting the marmintemp production.
	ExitMarmintemp(c *MarmintempContext)
}
