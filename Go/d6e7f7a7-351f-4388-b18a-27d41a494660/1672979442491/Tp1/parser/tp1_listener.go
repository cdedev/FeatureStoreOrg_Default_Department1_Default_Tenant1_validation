// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672979442491/Tp1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tp1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Tp1Listener is a complete listener for a parse tree produced by Tp1Parser.
type Tp1Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTp1 is called when entering the tp1 production.
	EnterTp1(c *Tp1Context)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTp1 is called when exiting the tp1 production.
	ExitTp1(c *Tp1Context)
}
