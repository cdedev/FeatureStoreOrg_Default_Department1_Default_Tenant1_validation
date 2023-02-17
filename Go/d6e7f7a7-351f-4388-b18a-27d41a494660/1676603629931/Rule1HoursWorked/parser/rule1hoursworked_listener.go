// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676603629931/Rule1HoursWorked.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rule1HoursWorked

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Rule1HoursWorkedListener is a complete listener for a parse tree produced by Rule1HoursWorkedParser.
type Rule1HoursWorkedListener interface {
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
