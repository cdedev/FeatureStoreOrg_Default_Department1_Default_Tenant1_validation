// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674837028967/Medal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Medal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMedalListener is a complete listener for a parse tree produced by MedalParser.
type BaseMedalListener struct{}

var _ MedalListener = &BaseMedalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMedalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMedalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMedalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMedalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMedalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMedalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMedal is called when production medal is entered.
func (s *BaseMedalListener) EnterMedal(ctx *MedalContext) {}

// ExitMedal is called when production medal is exited.
func (s *BaseMedalListener) ExitMedal(ctx *MedalContext) {}
