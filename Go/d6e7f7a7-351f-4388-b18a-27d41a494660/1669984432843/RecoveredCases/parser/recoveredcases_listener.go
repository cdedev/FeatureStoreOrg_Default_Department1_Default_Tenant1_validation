// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669984432843/RecoveredCases.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RecoveredCases

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RecoveredCasesListener is a complete listener for a parse tree produced by RecoveredCasesParser.
type RecoveredCasesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRecoveredcases is called when entering the recoveredcases production.
	EnterRecoveredcases(c *RecoveredcasesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRecoveredcases is called when exiting the recoveredcases production.
	ExitRecoveredcases(c *RecoveredcasesContext)
}
