// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669692269411/Diabetes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diabetes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiabetesListener is a complete listener for a parse tree produced by DiabetesParser.
type DiabetesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDiabetes is called when entering the diabetes production.
	EnterDiabetes(c *DiabetesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDiabetes is called when exiting the diabetes production.
	ExitDiabetes(c *DiabetesContext)
}
