// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690407274/Protein.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Protein

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ProteinListener is a complete listener for a parse tree produced by ProteinParser.
type ProteinListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterProtein is called when entering the protein production.
	EnterProtein(c *ProteinContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitProtein is called when exiting the protein production.
	ExitProtein(c *ProteinContext)
}
