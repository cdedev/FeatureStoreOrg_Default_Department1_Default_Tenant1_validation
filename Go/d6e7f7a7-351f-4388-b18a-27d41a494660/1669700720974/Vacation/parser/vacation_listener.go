// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700720974/Vacation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vacation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VacationListener is a complete listener for a parse tree produced by VacationParser.
type VacationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVacation is called when entering the vacation production.
	EnterVacation(c *VacationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVacation is called when exiting the vacation production.
	ExitVacation(c *VacationContext)
}
