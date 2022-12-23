// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671771923584/SessionDuration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SessionDuration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSessionDurationListener is a complete listener for a parse tree produced by SessionDurationParser.
type BaseSessionDurationListener struct{}

var _ SessionDurationListener = &BaseSessionDurationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSessionDurationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSessionDurationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSessionDurationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSessionDurationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSessionDurationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSessionDurationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSessionduration is called when production sessionduration is entered.
func (s *BaseSessionDurationListener) EnterSessionduration(ctx *SessiondurationContext) {}

// ExitSessionduration is called when production sessionduration is exited.
func (s *BaseSessionDurationListener) ExitSessionduration(ctx *SessiondurationContext) {}
