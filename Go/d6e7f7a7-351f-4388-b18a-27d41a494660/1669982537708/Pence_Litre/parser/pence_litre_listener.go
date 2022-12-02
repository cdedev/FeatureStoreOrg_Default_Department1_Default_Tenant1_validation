// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669982537708/Pence_Litre.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pence_Litre

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Pence_LitreListener is a complete listener for a parse tree produced by Pence_LitreParser.
type Pence_LitreListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPence_litre is called when entering the pence_litre production.
	EnterPence_litre(c *Pence_litreContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPence_litre is called when exiting the pence_litre production.
	ExitPence_litre(c *Pence_litreContext)
}
