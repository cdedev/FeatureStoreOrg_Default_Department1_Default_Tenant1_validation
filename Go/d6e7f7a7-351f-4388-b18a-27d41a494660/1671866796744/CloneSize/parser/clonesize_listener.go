// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671866796744/CloneSize.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CloneSize

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CloneSizeListener is a complete listener for a parse tree produced by CloneSizeParser.
type CloneSizeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterClonesize is called when entering the clonesize production.
	EnterClonesize(c *ClonesizeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitClonesize is called when exiting the clonesize production.
	ExitClonesize(c *ClonesizeContext)
}
