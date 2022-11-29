// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690358989/Energy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Energy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEnergyListener is a complete listener for a parse tree produced by EnergyParser.
type BaseEnergyListener struct{}

var _ EnergyListener = &BaseEnergyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEnergyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEnergyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEnergyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEnergyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEnergyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEnergyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEnergy is called when production energy is entered.
func (s *BaseEnergyListener) EnterEnergy(ctx *EnergyContext) {}

// ExitEnergy is called when production energy is exited.
func (s *BaseEnergyListener) ExitEnergy(ctx *EnergyContext) {}
