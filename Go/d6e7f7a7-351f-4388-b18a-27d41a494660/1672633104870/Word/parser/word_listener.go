// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672633104870/Word.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Word

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WordListener is a complete listener for a parse tree produced by WordParser.
type WordListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWord is called when entering the word production.
	EnterWord(c *WordContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWord is called when exiting the word production.
	ExitWord(c *WordContext)
}
