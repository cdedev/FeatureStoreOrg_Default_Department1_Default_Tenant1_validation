// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671602601656/Volume.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Volume

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VolumeListener is a complete listener for a parse tree produced by VolumeParser.
type VolumeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVolume is called when entering the volume production.
	EnterVolume(c *VolumeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVolume is called when exiting the volume production.
	ExitVolume(c *VolumeContext)
}
