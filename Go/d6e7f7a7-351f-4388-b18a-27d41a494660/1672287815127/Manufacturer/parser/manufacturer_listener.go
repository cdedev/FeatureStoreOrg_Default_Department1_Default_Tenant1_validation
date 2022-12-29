// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672287815127/Manufacturer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Manufacturer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ManufacturerListener is a complete listener for a parse tree produced by ManufacturerParser.
type ManufacturerListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterManufacturer is called when entering the manufacturer production.
	EnterManufacturer(c *ManufacturerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitManufacturer is called when exiting the manufacturer production.
	ExitManufacturer(c *ManufacturerContext)
}
