// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671601503289/AmountOfMoney.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AmountOfMoney

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AmountOfMoneyListener is a complete listener for a parse tree produced by AmountOfMoneyParser.
type AmountOfMoneyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAmountofmoney is called when entering the amountofmoney production.
	EnterAmountofmoney(c *AmountofmoneyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAmountofmoney is called when exiting the amountofmoney production.
	ExitAmountofmoney(c *AmountofmoneyContext)
}
