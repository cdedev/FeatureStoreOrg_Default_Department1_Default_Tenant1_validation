// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676462327753/WorkingCapacity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WorkingCapacity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WorkingCapacityListener is a complete listener for a parse tree produced by WorkingCapacityParser.
type WorkingCapacityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWorkingcapacity is called when entering the workingcapacity production.
	EnterWorkingcapacity(c *WorkingcapacityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWorkingcapacity is called when exiting the workingcapacity production.
	ExitWorkingcapacity(c *WorkingcapacityContext)
}
