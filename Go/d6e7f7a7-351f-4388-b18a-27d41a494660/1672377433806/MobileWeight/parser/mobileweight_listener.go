// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377433806/MobileWeight.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MobileWeight

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MobileWeightListener is a complete listener for a parse tree produced by MobileWeightParser.
type MobileWeightListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMobileweight is called when entering the mobileweight production.
	EnterMobileweight(c *MobileweightContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMobileweight is called when exiting the mobileweight production.
	ExitMobileweight(c *MobileweightContext)
}
