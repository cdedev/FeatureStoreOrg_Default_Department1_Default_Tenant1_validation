// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675937987570/Bandwidth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bandwidth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BandwidthListener is a complete listener for a parse tree produced by BandwidthParser.
type BandwidthListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBandwidth is called when entering the bandwidth production.
	EnterBandwidth(c *BandwidthContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBandwidth is called when exiting the bandwidth production.
	ExitBandwidth(c *BandwidthContext)
}
