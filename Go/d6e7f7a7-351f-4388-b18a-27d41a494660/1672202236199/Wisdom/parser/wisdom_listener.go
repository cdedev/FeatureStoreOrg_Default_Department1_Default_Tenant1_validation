// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202236199/Wisdom.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wisdom

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WisdomListener is a complete listener for a parse tree produced by WisdomParser.
type WisdomListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWisdom is called when entering the wisdom production.
	EnterWisdom(c *WisdomContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWisdom is called when exiting the wisdom production.
	ExitWisdom(c *WisdomContext)
}
