// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674793625076/Section.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Section

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SectionListener is a complete listener for a parse tree produced by SectionParser.
type SectionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSection is called when entering the section production.
	EnterSection(c *SectionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)
}
