// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675702851794/Keywords.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Keywords

import "github.com/antlr/antlr4/runtime/Go/antlr"

// KeywordsListener is a complete listener for a parse tree produced by KeywordsParser.
type KeywordsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterKeywords is called when entering the keywords production.
	EnterKeywords(c *KeywordsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitKeywords is called when exiting the keywords production.
	ExitKeywords(c *KeywordsContext)
}
