// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674790494867/Mass.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mass

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMassListener is a complete listener for a parse tree produced by MassParser.
type BaseMassListener struct{}

var _ MassListener = &BaseMassListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMassListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMassListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMassListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMassListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMassListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMassListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMass is called when production mass is entered.
func (s *BaseMassListener) EnterMass(ctx *MassContext) {}

// ExitMass is called when production mass is exited.
func (s *BaseMassListener) ExitMass(ctx *MassContext) {}
