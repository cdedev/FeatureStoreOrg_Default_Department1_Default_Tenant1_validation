// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624848798/District.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // District

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DistrictListener is a complete listener for a parse tree produced by DistrictParser.
type DistrictListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDistrict is called when entering the district production.
	EnterDistrict(c *DistrictContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDistrict is called when exiting the district production.
	ExitDistrict(c *DistrictContext)
}
