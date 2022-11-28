// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669659170926/SquareFeet.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SquareFeet

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SquareFeetListener is a complete listener for a parse tree produced by SquareFeetParser.
type SquareFeetListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSquarefeet is called when entering the squarefeet production.
	EnterSquarefeet(c *SquarefeetContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSquarefeet is called when exiting the squarefeet production.
	ExitSquarefeet(c *SquarefeetContext)
}
