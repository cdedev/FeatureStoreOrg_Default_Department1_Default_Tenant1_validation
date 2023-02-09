// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675932748157/CompanyType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CompanyType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CompanyTypeListener is a complete listener for a parse tree produced by CompanyTypeParser.
type CompanyTypeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCompanytype is called when entering the companytype production.
	EnterCompanytype(c *CompanytypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCompanytype is called when exiting the companytype production.
	ExitCompanytype(c *CompanytypeContext)
}
