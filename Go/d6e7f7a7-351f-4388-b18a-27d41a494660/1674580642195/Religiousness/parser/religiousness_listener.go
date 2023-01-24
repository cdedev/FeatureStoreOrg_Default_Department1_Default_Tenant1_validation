// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674580642195/Religiousness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Religiousness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ReligiousnessListener is a complete listener for a parse tree produced by ReligiousnessParser.
type ReligiousnessListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterReligiousness is called when entering the religiousness production.
	EnterReligiousness(c *ReligiousnessContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitReligiousness is called when exiting the religiousness production.
	ExitReligiousness(c *ReligiousnessContext)
}
