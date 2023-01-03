// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672717107908/Company.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Company

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CompanyListener is a complete listener for a parse tree produced by CompanyParser.
type CompanyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCompany is called when entering the company production.
	EnterCompany(c *CompanyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCompany is called when exiting the company production.
	ExitCompany(c *CompanyContext)
}
