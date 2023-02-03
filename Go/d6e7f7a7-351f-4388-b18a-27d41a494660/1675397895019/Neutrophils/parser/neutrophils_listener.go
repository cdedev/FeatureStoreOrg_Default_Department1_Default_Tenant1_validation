// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675397895019/Neutrophils.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Neutrophils

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NeutrophilsListener is a complete listener for a parse tree produced by NeutrophilsParser.
type NeutrophilsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNeutrophils is called when entering the neutrophils production.
	EnterNeutrophils(c *NeutrophilsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNeutrophils is called when exiting the neutrophils production.
	ExitNeutrophils(c *NeutrophilsContext)
}
