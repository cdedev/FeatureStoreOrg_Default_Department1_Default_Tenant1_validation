// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878395315/Photophobia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Photophobia

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PhotophobiaListener is a complete listener for a parse tree produced by PhotophobiaParser.
type PhotophobiaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPhotophobia is called when entering the photophobia production.
	EnterPhotophobia(c *PhotophobiaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPhotophobia is called when exiting the photophobia production.
	ExitPhotophobia(c *PhotophobiaContext)
}
