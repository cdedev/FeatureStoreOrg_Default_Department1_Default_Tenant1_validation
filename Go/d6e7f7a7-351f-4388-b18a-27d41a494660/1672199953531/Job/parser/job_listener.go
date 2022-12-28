// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672199953531/Job.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Job

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JobListener is a complete listener for a parse tree produced by JobParser.
type JobListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterJob is called when entering the job production.
	EnterJob(c *JobContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitJob is called when exiting the job production.
	ExitJob(c *JobContext)
}
