// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604039688/Bromine.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bromine

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BromineListener is a complete listener for a parse tree produced by BromineParser.
type BromineListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBromine is called when entering the bromine production.
	EnterBromine(c *BromineContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBromine is called when exiting the bromine production.
	ExitBromine(c *BromineContext)
}
