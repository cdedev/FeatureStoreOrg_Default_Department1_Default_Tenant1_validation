// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674096609856/Azimuth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Azimuth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AzimuthListener is a complete listener for a parse tree produced by AzimuthParser.
type AzimuthListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAzimuth is called when entering the azimuth production.
	EnterAzimuth(c *AzimuthContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAzimuth is called when exiting the azimuth production.
	ExitAzimuth(c *AzimuthContext)
}
