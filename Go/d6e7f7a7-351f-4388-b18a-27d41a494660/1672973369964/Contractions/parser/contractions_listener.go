// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672973369964/Contractions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Contractions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ContractionsListener is a complete listener for a parse tree produced by ContractionsParser.
type ContractionsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterContractions is called when entering the contractions production.
	EnterContractions(c *ContractionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitContractions is called when exiting the contractions production.
	ExitContractions(c *ContractionsContext)
}
