// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673926123247/Venomous.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Venomous

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VenomousListener is a complete listener for a parse tree produced by VenomousParser.
type VenomousListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVenomous is called when entering the venomous production.
	EnterVenomous(c *VenomousContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVenomous is called when exiting the venomous production.
	ExitVenomous(c *VenomousContext)
}
