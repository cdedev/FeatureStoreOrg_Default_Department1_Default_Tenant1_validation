// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669984432843/RecoveredCases.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RecoveredCases

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRecoveredCasesListener is a complete listener for a parse tree produced by RecoveredCasesParser.
type BaseRecoveredCasesListener struct{}

var _ RecoveredCasesListener = &BaseRecoveredCasesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRecoveredCasesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRecoveredCasesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRecoveredCasesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRecoveredCasesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRecoveredCasesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRecoveredCasesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRecoveredcases is called when production recoveredcases is entered.
func (s *BaseRecoveredCasesListener) EnterRecoveredcases(ctx *RecoveredcasesContext) {}

// ExitRecoveredcases is called when production recoveredcases is exited.
func (s *BaseRecoveredCasesListener) ExitRecoveredcases(ctx *RecoveredcasesContext) {}
