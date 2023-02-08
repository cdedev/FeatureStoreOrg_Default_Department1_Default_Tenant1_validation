// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675834395460/Sector.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sector

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SectorListener is a complete listener for a parse tree produced by SectorParser.
type SectorListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSector is called when entering the sector production.
	EnterSector(c *SectorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSector is called when exiting the sector production.
	ExitSector(c *SectorContext)
}
