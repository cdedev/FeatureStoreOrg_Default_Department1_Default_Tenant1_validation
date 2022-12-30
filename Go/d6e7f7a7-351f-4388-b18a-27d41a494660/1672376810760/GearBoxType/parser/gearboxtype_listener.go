// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376810760/GearBoxType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GearBoxType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GearBoxTypeListener is a complete listener for a parse tree produced by GearBoxTypeParser.
type GearBoxTypeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGearboxtype is called when entering the gearboxtype production.
	EnterGearboxtype(c *GearboxtypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGearboxtype is called when exiting the gearboxtype production.
	ExitGearboxtype(c *GearboxtypeContext)
}
