// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674096751985/InternetAccess.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // InternetAccess

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InternetAccessListener is a complete listener for a parse tree produced by InternetAccessParser.
type InternetAccessListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInternetaccess is called when entering the internetaccess production.
	EnterInternetaccess(c *InternetaccessContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInternetaccess is called when exiting the internetaccess production.
	ExitInternetaccess(c *InternetaccessContext)
}
