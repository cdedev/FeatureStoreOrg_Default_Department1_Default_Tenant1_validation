// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671695106646/Outdoor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Outdoor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOutdoorListener is a complete listener for a parse tree produced by OutdoorParser.
type BaseOutdoorListener struct{}

var _ OutdoorListener = &BaseOutdoorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOutdoorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOutdoorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOutdoorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOutdoorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOutdoorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOutdoorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOutdoor is called when production outdoor is entered.
func (s *BaseOutdoorListener) EnterOutdoor(ctx *OutdoorContext) {}

// ExitOutdoor is called when production outdoor is exited.
func (s *BaseOutdoorListener) ExitOutdoor(ctx *OutdoorContext) {}
