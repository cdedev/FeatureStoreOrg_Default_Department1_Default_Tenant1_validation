// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675182928861/SmokingStatus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SmokingStatus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSmokingStatusListener is a complete listener for a parse tree produced by SmokingStatusParser.
type BaseSmokingStatusListener struct{}

var _ SmokingStatusListener = &BaseSmokingStatusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSmokingStatusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSmokingStatusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSmokingStatusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSmokingStatusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSmokingStatusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSmokingStatusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSmokingstatus is called when production smokingstatus is entered.
func (s *BaseSmokingStatusListener) EnterSmokingstatus(ctx *SmokingstatusContext) {}

// ExitSmokingstatus is called when production smokingstatus is exited.
func (s *BaseSmokingStatusListener) ExitSmokingstatus(ctx *SmokingstatusContext) {}
