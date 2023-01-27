// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674839818224/Joblost.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Joblost

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JoblostListener is a complete listener for a parse tree produced by JoblostParser.
type JoblostListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterJoblost is called when entering the joblost production.
	EnterJoblost(c *JoblostContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitJoblost is called when exiting the joblost production.
	ExitJoblost(c *JoblostContext)
}
