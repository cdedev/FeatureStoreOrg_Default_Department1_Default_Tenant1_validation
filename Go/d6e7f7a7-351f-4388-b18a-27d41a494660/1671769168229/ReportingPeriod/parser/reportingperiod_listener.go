// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671769168229/ReportingPeriod.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ReportingPeriod

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ReportingPeriodListener is a complete listener for a parse tree produced by ReportingPeriodParser.
type ReportingPeriodListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterReportingperiod is called when entering the reportingperiod production.
	EnterReportingperiod(c *ReportingperiodContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitReportingperiod is called when exiting the reportingperiod production.
	ExitReportingperiod(c *ReportingperiodContext)
}
