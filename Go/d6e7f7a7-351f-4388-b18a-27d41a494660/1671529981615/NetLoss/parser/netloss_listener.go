// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671529981615/NetLoss.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NetLoss

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NetLossListener is a complete listener for a parse tree produced by NetLossParser.
type NetLossListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNetloss is called when entering the netloss production.
	EnterNetloss(c *NetlossContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNetloss is called when exiting the netloss production.
	ExitNetloss(c *NetlossContext)
}
