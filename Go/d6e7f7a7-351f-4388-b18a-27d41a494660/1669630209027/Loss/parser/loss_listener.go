// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669630209027/Loss.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Loss

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LossListener is a complete listener for a parse tree produced by LossParser.
type LossListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLoss is called when entering the loss production.
	EnterLoss(c *LossContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLoss is called when exiting the loss production.
	ExitLoss(c *LossContext)
}
