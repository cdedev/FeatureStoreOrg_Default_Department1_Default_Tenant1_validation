// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671698626799/JobLevel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JobLevel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJobLevelListener is a complete listener for a parse tree produced by JobLevelParser.
type BaseJobLevelListener struct{}

var _ JobLevelListener = &BaseJobLevelListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJobLevelListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJobLevelListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJobLevelListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJobLevelListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJobLevelListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJobLevelListener) ExitExpression(ctx *ExpressionContext) {}

// EnterJoblevel is called when production joblevel is entered.
func (s *BaseJobLevelListener) EnterJoblevel(ctx *JoblevelContext) {}

// ExitJoblevel is called when production joblevel is exited.
func (s *BaseJobLevelListener) ExitJoblevel(ctx *JoblevelContext) {}
