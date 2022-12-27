// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121452438/Dogs.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dogs

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DogsListener is a complete listener for a parse tree produced by DogsParser.
type DogsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDogs is called when entering the dogs production.
	EnterDogs(c *DogsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDogs is called when exiting the dogs production.
	ExitDogs(c *DogsContext)
}
