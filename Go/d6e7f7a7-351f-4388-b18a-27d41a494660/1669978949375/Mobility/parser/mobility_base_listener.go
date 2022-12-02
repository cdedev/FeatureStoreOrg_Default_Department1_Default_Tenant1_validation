// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669978949375/Mobility.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mobility

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMobilityListener is a complete listener for a parse tree produced by MobilityParser.
type BaseMobilityListener struct{}

var _ MobilityListener = &BaseMobilityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMobilityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMobilityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMobilityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMobilityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMobilityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMobilityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMobility is called when production mobility is entered.
func (s *BaseMobilityListener) EnterMobility(ctx *MobilityContext) {}

// ExitMobility is called when production mobility is exited.
func (s *BaseMobilityListener) ExitMobility(ctx *MobilityContext) {}
