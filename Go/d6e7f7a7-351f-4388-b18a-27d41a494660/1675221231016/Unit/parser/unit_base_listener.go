// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675221231016/Unit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Unit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseUnitListener is a complete listener for a parse tree produced by UnitParser.
type BaseUnitListener struct{}

var _ UnitListener = &BaseUnitListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUnitListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUnitListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUnitListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUnitListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseUnitListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseUnitListener) ExitExpression(ctx *ExpressionContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseUnitListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseUnitListener) ExitUnit(ctx *UnitContext) {}
