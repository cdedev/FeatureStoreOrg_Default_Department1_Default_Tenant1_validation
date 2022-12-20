// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671515379219/VehicleType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // VehicleType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VehicleTypeListener is a complete listener for a parse tree produced by VehicleTypeParser.
type VehicleTypeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVehicletype is called when entering the vehicletype production.
	EnterVehicletype(c *VehicletypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVehicletype is called when exiting the vehicletype production.
	ExitVehicletype(c *VehicletypeContext)
}
