// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669650939823/Pages.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pages

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PagesListener is a complete listener for a parse tree produced by PagesParser.
type PagesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPages is called when entering the pages production.
	EnterPages(c *PagesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPages is called when exiting the pages production.
	ExitPages(c *PagesContext)
}
