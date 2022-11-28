// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669651221577/Title.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Title

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TitleListener is a complete listener for a parse tree produced by TitleParser.
type TitleListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTitle is called when entering the title production.
	EnterTitle(c *TitleContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTitle is called when exiting the title production.
	ExitTitle(c *TitleContext)
}
