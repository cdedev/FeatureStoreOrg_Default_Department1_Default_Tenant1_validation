// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690110344/AddedSugar.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AddedSugar

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AddedSugarListener is a complete listener for a parse tree produced by AddedSugarParser.
type AddedSugarListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAddedsugar is called when entering the addedsugar production.
	EnterAddedsugar(c *AddedsugarContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAddedsugar is called when exiting the addedsugar production.
	ExitAddedsugar(c *AddedsugarContext)
}
