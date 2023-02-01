// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675223939602/Ethnicity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ethnicity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EthnicityListener is a complete listener for a parse tree produced by EthnicityParser.
type EthnicityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEthnicity is called when entering the ethnicity production.
	EnterEthnicity(c *EthnicityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEthnicity is called when exiting the ethnicity production.
	ExitEthnicity(c *EthnicityContext)
}
