// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076891972/FebMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FebMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FebMaxtempListener is a complete listener for a parse tree produced by FebMaxtempParser.
type FebMaxtempListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFebmaxtemp is called when entering the febmaxtemp production.
	EnterFebmaxtemp(c *FebmaxtempContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFebmaxtemp is called when exiting the febmaxtemp production.
	ExitFebmaxtemp(c *FebmaxtempContext)
}
