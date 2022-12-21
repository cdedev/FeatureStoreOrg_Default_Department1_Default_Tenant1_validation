// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671595271782/Mileage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mileage

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MileageListener is a complete listener for a parse tree produced by MileageParser.
type MileageListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMileage is called when entering the mileage production.
	EnterMileage(c *MileageContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMileage is called when exiting the mileage production.
	ExitMileage(c *MileageContext)
}