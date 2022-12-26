// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076932995/FebMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FebMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FebMintempListener is a complete listener for a parse tree produced by FebMintempParser.
type FebMintempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFebmintemp is called when entering the febmintemp production.
	EnterFebmintemp(c *FebmintempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFebmintemp is called when exiting the febmintemp production.
	ExitFebmintemp(c *FebmintempContext)
}
