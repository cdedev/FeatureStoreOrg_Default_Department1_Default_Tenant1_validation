// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669695485622/Moonphase.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Moonphase

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMoonphaseListener is a complete listener for a parse tree produced by MoonphaseParser.
type BaseMoonphaseListener struct{}

var _ MoonphaseListener = &BaseMoonphaseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMoonphaseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMoonphaseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMoonphaseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMoonphaseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMoonphaseListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMoonphaseListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMoonphase is called when production moonphase is entered.
func (s *BaseMoonphaseListener) EnterMoonphase(ctx *MoonphaseContext) {}

// ExitMoonphase is called when production moonphase is exited.
func (s *BaseMoonphaseListener) ExitMoonphase(ctx *MoonphaseContext) {}
