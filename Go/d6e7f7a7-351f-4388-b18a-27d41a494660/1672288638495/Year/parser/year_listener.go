// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672288638495/Year.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Year

import "github.com/antlr/antlr4/runtime/Go/antlr"

// YearListener is a complete listener for a parse tree produced by YearParser.
type YearListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterYear is called when entering the year production.
	EnterYear(c *YearContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitYear is called when exiting the year production.
	ExitYear(c *YearContext)
}
