// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671768897139/FuelConsumption.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FuelConsumption

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFuelConsumptionListener is a complete listener for a parse tree produced by FuelConsumptionParser.
type BaseFuelConsumptionListener struct{}

var _ FuelConsumptionListener = &BaseFuelConsumptionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFuelConsumptionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFuelConsumptionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFuelConsumptionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFuelConsumptionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFuelConsumptionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFuelConsumptionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFuelconsumption is called when production fuelconsumption is entered.
func (s *BaseFuelConsumptionListener) EnterFuelconsumption(ctx *FuelconsumptionContext) {}

// ExitFuelconsumption is called when production fuelconsumption is exited.
func (s *BaseFuelConsumptionListener) ExitFuelconsumption(ctx *FuelconsumptionContext) {}
