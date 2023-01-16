// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878144865/Dysphasia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dysphasia

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DysphasiaListener is a complete listener for a parse tree produced by DysphasiaParser.
type DysphasiaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDysphasia is called when entering the dysphasia production.
	EnterDysphasia(c *DysphasiaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDysphasia is called when exiting the dysphasia production.
	ExitDysphasia(c *DysphasiaContext)
}
