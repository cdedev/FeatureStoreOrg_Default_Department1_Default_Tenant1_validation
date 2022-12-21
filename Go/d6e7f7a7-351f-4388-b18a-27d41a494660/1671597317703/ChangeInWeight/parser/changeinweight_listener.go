// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671597317703/ChangeInWeight.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ChangeInWeight

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChangeInWeightListener is a complete listener for a parse tree produced by ChangeInWeightParser.
type ChangeInWeightListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterChangeinweight is called when entering the changeinweight production.
	EnterChangeinweight(c *ChangeinweightContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitChangeinweight is called when exiting the changeinweight production.
	ExitChangeinweight(c *ChangeinweightContext)
}
