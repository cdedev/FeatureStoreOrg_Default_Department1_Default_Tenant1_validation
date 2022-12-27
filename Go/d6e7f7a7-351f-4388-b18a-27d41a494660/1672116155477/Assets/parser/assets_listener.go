// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672116155477/Assets.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Assets

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AssetsListener is a complete listener for a parse tree produced by AssetsParser.
type AssetsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssets is called when entering the assets production.
	EnterAssets(c *AssetsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssets is called when exiting the assets production.
	ExitAssets(c *AssetsContext)
}
