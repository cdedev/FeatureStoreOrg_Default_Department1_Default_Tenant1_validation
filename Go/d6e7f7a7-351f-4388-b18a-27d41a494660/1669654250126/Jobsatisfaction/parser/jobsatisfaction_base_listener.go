// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669654250126/Jobsatisfaction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jobsatisfaction

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJobsatisfactionListener is a complete listener for a parse tree produced by JobsatisfactionParser.
type BaseJobsatisfactionListener struct{}

var _ JobsatisfactionListener = &BaseJobsatisfactionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJobsatisfactionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJobsatisfactionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJobsatisfactionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJobsatisfactionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJobsatisfactionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJobsatisfactionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterJobsatisfaction is called when production jobsatisfaction is entered.
func (s *BaseJobsatisfactionListener) EnterJobsatisfaction(ctx *JobsatisfactionContext) {}

// ExitJobsatisfaction is called when production jobsatisfaction is exited.
func (s *BaseJobsatisfactionListener) ExitJobsatisfaction(ctx *JobsatisfactionContext) {}
