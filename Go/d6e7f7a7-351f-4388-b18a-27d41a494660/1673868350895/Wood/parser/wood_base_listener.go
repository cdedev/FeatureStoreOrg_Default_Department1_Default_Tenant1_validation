// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673868350895/Wood.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wood

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWoodListener is a complete listener for a parse tree produced by WoodParser.
type BaseWoodListener struct{}

var _ WoodListener = &BaseWoodListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWoodListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWoodListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWoodListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWoodListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWoodListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWoodListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWood is called when production wood is entered.
func (s *BaseWoodListener) EnterWood(ctx *WoodContext) {}

// ExitWood is called when production wood is exited.
func (s *BaseWoodListener) ExitWood(ctx *WoodContext) {}
