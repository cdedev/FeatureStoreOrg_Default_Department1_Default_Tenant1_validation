// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121133708/NonHumanPrimates.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NonHumanPrimates

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NonHumanPrimatesListener is a complete listener for a parse tree produced by NonHumanPrimatesParser.
type NonHumanPrimatesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNonhumanprimates is called when entering the nonhumanprimates production.
	EnterNonhumanprimates(c *NonhumanprimatesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNonhumanprimates is called when exiting the nonhumanprimates production.
	ExitNonhumanprimates(c *NonhumanprimatesContext)
}
