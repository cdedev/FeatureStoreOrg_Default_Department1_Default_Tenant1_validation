// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377534929/FrontCamera.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FrontCamera

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FrontCameraListener is a complete listener for a parse tree produced by FrontCameraParser.
type FrontCameraListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFrontcamera is called when entering the frontcamera production.
	EnterFrontcamera(c *FrontcameraContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFrontcamera is called when exiting the frontcamera production.
	ExitFrontcamera(c *FrontcameraContext)
}
