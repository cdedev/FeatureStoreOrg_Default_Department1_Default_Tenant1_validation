// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377576745/MobileDep.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MobileDep

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MobileDepListener is a complete listener for a parse tree produced by MobileDepParser.
type MobileDepListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMobiledep is called when entering the mobiledep production.
	EnterMobiledep(c *MobiledepContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMobiledep is called when exiting the mobiledep production.
	ExitMobiledep(c *MobiledepContext)
}
