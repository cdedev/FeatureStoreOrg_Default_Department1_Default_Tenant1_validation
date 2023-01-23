// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674455214945/Treatment.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Treatment

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TreatmentListener is a complete listener for a parse tree produced by TreatmentParser.
type TreatmentListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTreatment is called when entering the treatment production.
	EnterTreatment(c *TreatmentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTreatment is called when exiting the treatment production.
	ExitTreatment(c *TreatmentContext)
}
