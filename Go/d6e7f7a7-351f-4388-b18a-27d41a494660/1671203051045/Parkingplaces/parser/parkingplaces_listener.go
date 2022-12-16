// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671203051045/Parkingplaces.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Parkingplaces

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ParkingplacesListener is a complete listener for a parse tree produced by ParkingplacesParser.
type ParkingplacesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterParkingplaces is called when entering the parkingplaces production.
	EnterParkingplaces(c *ParkingplacesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitParkingplaces is called when exiting the parkingplaces production.
	ExitParkingplaces(c *ParkingplacesContext)
}
