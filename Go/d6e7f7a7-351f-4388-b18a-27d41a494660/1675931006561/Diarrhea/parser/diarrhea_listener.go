// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675931006561/Diarrhea.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diarrhea

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiarrheaListener is a complete listener for a parse tree produced by DiarrheaParser.
type DiarrheaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDiarrhea is called when entering the diarrhea production.
	EnterDiarrhea(c *DiarrheaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDiarrhea is called when exiting the diarrhea production.
	ExitDiarrhea(c *DiarrheaContext)
}
