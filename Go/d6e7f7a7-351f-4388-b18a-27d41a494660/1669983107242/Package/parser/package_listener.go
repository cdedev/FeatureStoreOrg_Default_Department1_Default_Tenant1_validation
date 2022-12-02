// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669983107242/Package.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Package

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PackageListener is a complete listener for a parse tree produced by PackageParser.
type PackageListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPackage is called when entering the package production.
	EnterPackage(c *PackageContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPackage is called when exiting the package production.
	ExitPackage(c *PackageContext)
}
