// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675224395957/License.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // License

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LicenseListener is a complete listener for a parse tree produced by LicenseParser.
type LicenseListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLicense is called when entering the license production.
	EnterLicense(c *LicenseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLicense is called when exiting the license production.
	ExitLicense(c *LicenseContext)
}
