// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377226996/BatteryPower.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BatteryPower

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBatteryPowerListener is a complete listener for a parse tree produced by BatteryPowerParser.
type BaseBatteryPowerListener struct{}

var _ BatteryPowerListener = &BaseBatteryPowerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBatteryPowerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBatteryPowerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBatteryPowerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBatteryPowerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBatteryPowerListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBatteryPowerListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBatterypower is called when production batterypower is entered.
func (s *BaseBatteryPowerListener) EnterBatterypower(ctx *BatterypowerContext) {}

// ExitBatterypower is called when production batterypower is exited.
func (s *BaseBatteryPowerListener) ExitBatterypower(ctx *BatterypowerContext) {}
