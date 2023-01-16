// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673867910307/Paper.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Paper

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PaperListener is a complete listener for a parse tree produced by PaperParser.
type PaperListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPaper is called when entering the paper production.
	EnterPaper(c *PaperContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPaper is called when exiting the paper production.
	ExitPaper(c *PaperContext)
}
