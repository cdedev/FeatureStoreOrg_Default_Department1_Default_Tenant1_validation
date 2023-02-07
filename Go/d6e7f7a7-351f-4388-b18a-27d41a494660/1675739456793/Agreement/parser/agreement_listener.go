// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675739456793/Agreement.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Agreement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AgreementListener is a complete listener for a parse tree produced by AgreementParser.
type AgreementListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAgreement is called when entering the agreement production.
	EnterAgreement(c *AgreementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAgreement is called when exiting the agreement production.
	ExitAgreement(c *AgreementContext)
}
