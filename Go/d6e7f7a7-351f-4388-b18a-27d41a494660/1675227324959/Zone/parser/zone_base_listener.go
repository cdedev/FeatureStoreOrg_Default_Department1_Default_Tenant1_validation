// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675227324959/Zone.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Zone

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseZoneListener is a complete listener for a parse tree produced by ZoneParser.
type BaseZoneListener struct{}

var _ ZoneListener = &BaseZoneListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseZoneListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseZoneListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseZoneListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseZoneListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseZoneListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseZoneListener) ExitExpression(ctx *ExpressionContext) {}

// EnterZone is called when production zone is entered.
func (s *BaseZoneListener) EnterZone(ctx *ZoneContext) {}

// ExitZone is called when production zone is exited.
func (s *BaseZoneListener) ExitZone(ctx *ZoneContext) {}
