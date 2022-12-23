// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671767294086/Killed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Killed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseKilledListener is a complete listener for a parse tree produced by KilledParser.
type BaseKilledListener struct{}

var _ KilledListener = &BaseKilledListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKilledListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKilledListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKilledListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKilledListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseKilledListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseKilledListener) ExitExpression(ctx *ExpressionContext) {}

// EnterKilled is called when production killed is entered.
func (s *BaseKilledListener) EnterKilled(ctx *KilledContext) {}

// ExitKilled is called when production killed is exited.
func (s *BaseKilledListener) ExitKilled(ctx *KilledContext) {}
