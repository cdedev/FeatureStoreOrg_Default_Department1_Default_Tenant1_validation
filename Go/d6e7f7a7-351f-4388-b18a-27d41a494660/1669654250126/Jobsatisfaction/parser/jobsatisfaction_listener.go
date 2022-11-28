// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669654250126/Jobsatisfaction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jobsatisfaction

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JobsatisfactionListener is a complete listener for a parse tree produced by JobsatisfactionParser.
type JobsatisfactionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterJobsatisfaction is called when entering the jobsatisfaction production.
	EnterJobsatisfaction(c *JobsatisfactionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitJobsatisfaction is called when exiting the jobsatisfaction production.
	ExitJobsatisfaction(c *JobsatisfactionContext)
}
