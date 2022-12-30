// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672372092240/Code.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Code

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CodeListener is a complete listener for a parse tree produced by CodeParser.
type CodeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCode is called when entering the code production.
	EnterCode(c *CodeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCode is called when exiting the code production.
	ExitCode(c *CodeContext)
}
