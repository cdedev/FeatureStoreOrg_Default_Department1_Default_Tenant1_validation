// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690427235/MonthAsCustomer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MonthAsCustomer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MonthAsCustomerListener is a complete listener for a parse tree produced by MonthAsCustomerParser.
type MonthAsCustomerListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMonthascustomer is called when entering the monthascustomer production.
	EnterMonthascustomer(c *MonthascustomerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMonthascustomer is called when exiting the monthascustomer production.
	ExitMonthascustomer(c *MonthascustomerContext)
}
