// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672199953531/Job.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Job

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJobListener is a complete listener for a parse tree produced by JobParser.
type BaseJobListener struct{}

var _ JobListener = &BaseJobListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJobListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJobListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJobListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJobListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJobListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJobListener) ExitExpression(ctx *ExpressionContext) {}

// EnterJob is called when production job is entered.
func (s *BaseJobListener) EnterJob(ctx *JobContext) {}

// ExitJob is called when production job is exited.
func (s *BaseJobListener) ExitJob(ctx *JobContext) {}
