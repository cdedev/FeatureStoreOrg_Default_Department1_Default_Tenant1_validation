// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672204141055/FederalOccupation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FederalOccupation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFederalOccupationListener is a complete listener for a parse tree produced by FederalOccupationParser.
type BaseFederalOccupationListener struct{}

var _ FederalOccupationListener = &BaseFederalOccupationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFederalOccupationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFederalOccupationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFederalOccupationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFederalOccupationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFederalOccupationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFederalOccupationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFederaloccupation is called when production federaloccupation is entered.
func (s *BaseFederalOccupationListener) EnterFederaloccupation(ctx *FederaloccupationContext) {}

// ExitFederaloccupation is called when production federaloccupation is exited.
func (s *BaseFederalOccupationListener) ExitFederaloccupation(ctx *FederaloccupationContext) {}
