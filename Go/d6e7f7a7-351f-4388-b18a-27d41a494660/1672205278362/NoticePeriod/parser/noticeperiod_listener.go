// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205278362/NoticePeriod.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NoticePeriod

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NoticePeriodListener is a complete listener for a parse tree produced by NoticePeriodParser.
type NoticePeriodListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNoticeperiod is called when entering the noticeperiod production.
	EnterNoticeperiod(c *NoticeperiodContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNoticeperiod is called when exiting the noticeperiod production.
	ExitNoticeperiod(c *NoticeperiodContext)
}
