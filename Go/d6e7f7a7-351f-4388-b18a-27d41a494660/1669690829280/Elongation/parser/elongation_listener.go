// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690829280/Elongation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Elongation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ElongationListener is a complete listener for a parse tree produced by ElongationParser.
type ElongationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterElongation is called when entering the elongation production.
	EnterElongation(c *ElongationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitElongation is called when exiting the elongation production.
	ExitElongation(c *ElongationContext)
}
