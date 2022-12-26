// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077184093/MayMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MayMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MayMintempListener is a complete listener for a parse tree produced by MayMintempParser.
type MayMintempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMaymintemp is called when entering the maymintemp production.
	EnterMaymintemp(c *MaymintempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMaymintemp is called when exiting the maymintemp production.
	ExitMaymintemp(c *MaymintempContext)
}
