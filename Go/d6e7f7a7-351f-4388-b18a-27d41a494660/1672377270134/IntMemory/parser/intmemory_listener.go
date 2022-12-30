// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377270134/IntMemory.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IntMemory

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IntMemoryListener is a complete listener for a parse tree produced by IntMemoryParser.
type IntMemoryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIntmemory is called when entering the intmemory production.
	EnterIntmemory(c *IntmemoryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIntmemory is called when exiting the intmemory production.
	ExitIntmemory(c *IntmemoryContext)
}
