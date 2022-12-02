// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669982272752/FuelType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FuelType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFuelTypeListener is a complete listener for a parse tree produced by FuelTypeParser.
type BaseFuelTypeListener struct{}

var _ FuelTypeListener = &BaseFuelTypeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFuelTypeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFuelTypeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFuelTypeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFuelTypeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFuelTypeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFuelTypeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFueltype is called when production fueltype is entered.
func (s *BaseFuelTypeListener) EnterFueltype(ctx *FueltypeContext) {}

// ExitFueltype is called when production fueltype is exited.
func (s *BaseFuelTypeListener) ExitFueltype(ctx *FueltypeContext) {}
