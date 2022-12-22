// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690299425/IncidentSeverity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IncidentSeverity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IncidentSeverityListener is a complete listener for a parse tree produced by IncidentSeverityParser.
type IncidentSeverityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIncidentseverity is called when entering the incidentseverity production.
	EnterIncidentseverity(c *IncidentseverityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIncidentseverity is called when exiting the incidentseverity production.
	ExitIncidentseverity(c *IncidentseverityContext)
}
