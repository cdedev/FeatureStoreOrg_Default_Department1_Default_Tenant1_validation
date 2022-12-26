// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672078015413/NovMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NovMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NovMintempListener is a complete listener for a parse tree produced by NovMintempParser.
type NovMintempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNovmintemp is called when entering the novmintemp production.
	EnterNovmintemp(c *NovmintempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNovmintemp is called when exiting the novmintemp production.
	ExitNovmintemp(c *NovmintempContext)
}
