// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604154500/Californium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Californium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CaliforniumListener is a complete listener for a parse tree produced by CaliforniumParser.
type CaliforniumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCalifornium is called when entering the californium production.
	EnterCalifornium(c *CaliforniumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCalifornium is called when exiting the californium production.
	ExitCalifornium(c *CaliforniumContext)
}
