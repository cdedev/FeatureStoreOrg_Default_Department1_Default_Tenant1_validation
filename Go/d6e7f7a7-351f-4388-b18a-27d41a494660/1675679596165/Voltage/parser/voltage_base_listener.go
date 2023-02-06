// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675679596165/Voltage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Voltage

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVoltageListener is a complete listener for a parse tree produced by VoltageParser.
type BaseVoltageListener struct{}

var _ VoltageListener = &BaseVoltageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVoltageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVoltageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVoltageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVoltageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVoltageListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVoltageListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVoltage is called when production voltage is entered.
func (s *BaseVoltageListener) EnterVoltage(ctx *VoltageContext) {}

// ExitVoltage is called when production voltage is exited.
func (s *BaseVoltageListener) ExitVoltage(ctx *VoltageContext) {}
