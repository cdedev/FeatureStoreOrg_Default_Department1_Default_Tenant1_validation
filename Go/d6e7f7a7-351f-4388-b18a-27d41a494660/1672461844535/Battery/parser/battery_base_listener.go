// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672461844535/Battery.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Battery

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBatteryListener is a complete listener for a parse tree produced by BatteryParser.
type BaseBatteryListener struct{}

var _ BatteryListener = &BaseBatteryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBatteryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBatteryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBatteryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBatteryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBatteryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBatteryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBattery is called when production battery is entered.
func (s *BaseBatteryListener) EnterBattery(ctx *BatteryContext) {}

// ExitBattery is called when production battery is exited.
func (s *BaseBatteryListener) ExitBattery(ctx *BatteryContext) {}
