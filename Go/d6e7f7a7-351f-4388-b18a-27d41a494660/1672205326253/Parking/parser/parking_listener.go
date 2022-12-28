// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205326253/Parking.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Parking

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ParkingListener is a complete listener for a parse tree produced by ParkingParser.
type ParkingListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterParking is called when entering the parking production.
	EnterParking(c *ParkingContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitParking is called when exiting the parking production.
	ExitParking(c *ParkingContext)
}
