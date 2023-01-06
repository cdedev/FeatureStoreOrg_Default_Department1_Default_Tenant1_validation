// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672979299463/Fm1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fm1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Fm1Listener is a complete listener for a parse tree produced by Fm1Parser.
type Fm1Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFm1 is called when entering the fm1 production.
	EnterFm1(c *Fm1Context)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFm1 is called when exiting the fm1 production.
	ExitFm1(c *Fm1Context)
}
