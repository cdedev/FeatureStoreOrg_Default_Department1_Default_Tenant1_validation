// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675742249049/CompanyName.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CompanyName

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CompanyNameListener is a complete listener for a parse tree produced by CompanyNameParser.
type CompanyNameListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCompanyname is called when entering the companyname production.
	EnterCompanyname(c *CompanynameContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCompanyname is called when exiting the companyname production.
	ExitCompanyname(c *CompanynameContext)
}
