// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671695074852/Indoor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Indoor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIndoorListener is a complete listener for a parse tree produced by IndoorParser.
type BaseIndoorListener struct{}

var _ IndoorListener = &BaseIndoorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIndoorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIndoorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIndoorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIndoorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIndoorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIndoorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIndoor is called when production indoor is entered.
func (s *BaseIndoorListener) EnterIndoor(ctx *IndoorContext) {}

// ExitIndoor is called when production indoor is exited.
func (s *BaseIndoorListener) ExitIndoor(ctx *IndoorContext) {}
