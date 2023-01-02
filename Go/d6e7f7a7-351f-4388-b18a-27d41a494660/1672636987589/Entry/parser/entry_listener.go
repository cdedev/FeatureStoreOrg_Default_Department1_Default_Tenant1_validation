// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672636987589/Entry.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Entry

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EntryListener is a complete listener for a parse tree produced by EntryParser.
type EntryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEntry is called when entering the entry production.
	EnterEntry(c *EntryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEntry is called when exiting the entry production.
	ExitEntry(c *EntryContext)
}
