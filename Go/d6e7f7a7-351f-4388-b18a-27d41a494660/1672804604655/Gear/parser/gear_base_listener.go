// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672804604655/Gear.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gear

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGearListener is a complete listener for a parse tree produced by GearParser.
type BaseGearListener struct{}

var _ GearListener = &BaseGearListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGearListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGearListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGearListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGearListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGearListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGearListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGear is called when production gear is entered.
func (s *BaseGearListener) EnterGear(ctx *GearContext) {}

// ExitGear is called when production gear is exited.
func (s *BaseGearListener) ExitGear(ctx *GearContext) {}
