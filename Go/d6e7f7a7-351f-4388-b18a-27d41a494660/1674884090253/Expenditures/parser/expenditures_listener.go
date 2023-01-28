// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674884090253/Expenditures.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expenditures

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExpendituresListener is a complete listener for a parse tree produced by ExpendituresParser.
type ExpendituresListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpenditures is called when entering the expenditures production.
	EnterExpenditures(c *ExpendituresContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpenditures is called when exiting the expenditures production.
	ExitExpenditures(c *ExpendituresContext)
}
