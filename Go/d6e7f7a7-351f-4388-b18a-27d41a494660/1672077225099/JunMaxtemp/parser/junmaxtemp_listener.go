// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077225099/JunMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JunMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JunMaxtempListener is a complete listener for a parse tree produced by JunMaxtempParser.
type JunMaxtempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterJunmaxtemp is called when entering the junmaxtemp production.
	EnterJunmaxtemp(c *JunmaxtempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitJunmaxtemp is called when exiting the junmaxtemp production.
	ExitJunmaxtemp(c *JunmaxtempContext)
}
