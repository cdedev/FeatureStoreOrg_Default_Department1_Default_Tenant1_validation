// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672985492567/Vendornoproduits.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vendornoproduits

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VendornoproduitsListener is a complete listener for a parse tree produced by VendornoproduitsParser.
type VendornoproduitsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVendornoproduits is called when entering the vendornoproduits production.
	EnterVendornoproduits(c *VendornoproduitsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVendornoproduits is called when exiting the vendornoproduits production.
	ExitVendornoproduits(c *VendornoproduitsContext)
}
