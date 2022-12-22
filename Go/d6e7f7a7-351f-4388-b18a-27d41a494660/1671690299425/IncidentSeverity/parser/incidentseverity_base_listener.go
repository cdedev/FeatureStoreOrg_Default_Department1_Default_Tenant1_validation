// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690299425/IncidentSeverity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IncidentSeverity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIncidentSeverityListener is a complete listener for a parse tree produced by IncidentSeverityParser.
type BaseIncidentSeverityListener struct{}

var _ IncidentSeverityListener = &BaseIncidentSeverityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIncidentSeverityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIncidentSeverityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIncidentSeverityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIncidentSeverityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIncidentSeverityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIncidentSeverityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIncidentseverity is called when production incidentseverity is entered.
func (s *BaseIncidentSeverityListener) EnterIncidentseverity(ctx *IncidentseverityContext) {}

// ExitIncidentseverity is called when production incidentseverity is exited.
func (s *BaseIncidentSeverityListener) ExitIncidentseverity(ctx *IncidentseverityContext) {}
