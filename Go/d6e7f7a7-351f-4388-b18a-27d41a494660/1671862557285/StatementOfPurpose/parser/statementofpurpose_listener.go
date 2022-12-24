// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671862557285/StatementOfPurpose.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // StatementOfPurpose

import "github.com/antlr/antlr4/runtime/Go/antlr"

// StatementOfPurposeListener is a complete listener for a parse tree produced by StatementOfPurposeParser.
type StatementOfPurposeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterStatementofpurpose is called when entering the statementofpurpose production.
	EnterStatementofpurpose(c *StatementofpurposeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitStatementofpurpose is called when exiting the statementofpurpose production.
	ExitStatementofpurpose(c *StatementofpurposeContext)
}
