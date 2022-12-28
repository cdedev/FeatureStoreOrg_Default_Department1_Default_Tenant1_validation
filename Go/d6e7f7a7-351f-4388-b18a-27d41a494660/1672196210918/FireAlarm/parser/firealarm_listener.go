// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672196210918/FireAlarm.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FireAlarm

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FireAlarmListener is a complete listener for a parse tree produced by FireAlarmParser.
type FireAlarmListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFirealarm is called when entering the firealarm production.
	EnterFirealarm(c *FirealarmContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFirealarm is called when exiting the firealarm production.
	ExitFirealarm(c *FirealarmContext)
}
