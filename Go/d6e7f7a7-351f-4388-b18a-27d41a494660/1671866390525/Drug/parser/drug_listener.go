// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671866390525/Drug.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Drug

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DrugListener is a complete listener for a parse tree produced by DrugParser.
type DrugListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDrug is called when entering the drug production.
	EnterDrug(c *DrugContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDrug is called when exiting the drug production.
	ExitDrug(c *DrugContext)
}
