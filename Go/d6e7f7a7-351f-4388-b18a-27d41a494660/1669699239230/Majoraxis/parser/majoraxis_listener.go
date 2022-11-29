// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669699239230/Majoraxis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Majoraxis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MajoraxisListener is a complete listener for a parse tree produced by MajoraxisParser.
type MajoraxisListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMajoraxis is called when entering the majoraxis production.
	EnterMajoraxis(c *MajoraxisContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMajoraxis is called when exiting the majoraxis production.
	ExitMajoraxis(c *MajoraxisContext)
}
