// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675231501618/HoursWorked.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // HoursWorked

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HoursWorkedListener is a complete listener for a parse tree produced by HoursWorkedParser.
type HoursWorkedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHoursworked is called when entering the hoursworked production.
	EnterHoursworked(c *HoursworkedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHoursworked is called when exiting the hoursworked production.
	ExitHoursworked(c *HoursworkedContext)
}
