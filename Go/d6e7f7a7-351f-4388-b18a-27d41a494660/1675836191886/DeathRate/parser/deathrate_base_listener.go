// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675836191886/DeathRate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DeathRate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDeathRateListener is a complete listener for a parse tree produced by DeathRateParser.
type BaseDeathRateListener struct{}

var _ DeathRateListener = &BaseDeathRateListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDeathRateListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDeathRateListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDeathRateListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDeathRateListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDeathRateListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDeathRateListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDeathrate is called when production deathrate is entered.
func (s *BaseDeathRateListener) EnterDeathrate(ctx *DeathrateContext) {}

// ExitDeathrate is called when production deathrate is exited.
func (s *BaseDeathRateListener) ExitDeathrate(ctx *DeathrateContext) {}
