// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674104198617/Cloudcover.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cloudcover

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CloudcoverListener is a complete listener for a parse tree produced by CloudcoverParser.
type CloudcoverListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCloudcover is called when entering the cloudcover production.
	EnterCloudcover(c *CloudcoverContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCloudcover is called when exiting the cloudcover production.
	ExitCloudcover(c *CloudcoverContext)
}
