// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672373090178/BPStable.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BPStable

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BPStableListener is a complete listener for a parse tree produced by BPStableParser.
type BPStableListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBpstable is called when entering the bpstable production.
	EnterBpstable(c *BpstableContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBpstable is called when exiting the bpstable production.
	ExitBpstable(c *BpstableContext)
}
