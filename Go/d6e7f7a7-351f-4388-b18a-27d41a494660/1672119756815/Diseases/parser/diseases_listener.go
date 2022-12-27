// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119756815/Diseases.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diseases

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiseasesListener is a complete listener for a parse tree produced by DiseasesParser.
type DiseasesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDiseases is called when entering the diseases production.
	EnterDiseases(c *DiseasesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDiseases is called when exiting the diseases production.
	ExitDiseases(c *DiseasesContext)
}
