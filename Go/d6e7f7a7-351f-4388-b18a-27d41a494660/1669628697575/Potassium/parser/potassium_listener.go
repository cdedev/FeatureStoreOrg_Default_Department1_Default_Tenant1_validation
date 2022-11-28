// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628697575/Potassium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Potassium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PotassiumListener is a complete listener for a parse tree produced by PotassiumParser.
type PotassiumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPotassium is called when entering the potassium production.
	EnterPotassium(c *PotassiumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPotassium is called when exiting the potassium production.
	ExitPotassium(c *PotassiumContext)
}
