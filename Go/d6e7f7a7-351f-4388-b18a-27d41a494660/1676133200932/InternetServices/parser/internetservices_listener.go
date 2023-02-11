// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676133200932/InternetServices.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // InternetServices

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InternetServicesListener is a complete listener for a parse tree produced by InternetServicesParser.
type InternetServicesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInternetservices is called when entering the internetservices production.
	EnterInternetservices(c *InternetservicesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInternetservices is called when exiting the internetservices production.
	ExitInternetservices(c *InternetservicesContext)
}
