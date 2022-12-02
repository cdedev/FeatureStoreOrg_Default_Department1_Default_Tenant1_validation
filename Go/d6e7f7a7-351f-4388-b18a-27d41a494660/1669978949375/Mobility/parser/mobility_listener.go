// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669978949375/Mobility.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mobility

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MobilityListener is a complete listener for a parse tree produced by MobilityParser.
type MobilityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMobility is called when entering the mobility production.
	EnterMobility(c *MobilityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMobility is called when exiting the mobility production.
	ExitMobility(c *MobilityContext)
}
