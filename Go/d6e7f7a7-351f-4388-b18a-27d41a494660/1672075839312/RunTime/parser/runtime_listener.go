// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672075839312/RunTime.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RunTime

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RunTimeListener is a complete listener for a parse tree produced by RunTimeParser.
type RunTimeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRuntime is called when entering the runtime production.
	EnterRuntime(c *RuntimeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRuntime is called when exiting the runtime production.
	ExitRuntime(c *RuntimeContext)
}
