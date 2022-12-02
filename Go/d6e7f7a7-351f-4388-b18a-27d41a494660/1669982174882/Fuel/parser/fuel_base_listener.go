// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669982174882/Fuel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fuel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFuelListener is a complete listener for a parse tree produced by FuelParser.
type BaseFuelListener struct{}

var _ FuelListener = &BaseFuelListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFuelListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFuelListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFuelListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFuelListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFuelListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFuelListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFuel is called when production fuel is entered.
func (s *BaseFuelListener) EnterFuel(ctx *FuelContext) {}

// ExitFuel is called when production fuel is exited.
func (s *BaseFuelListener) ExitFuel(ctx *FuelContext) {}
