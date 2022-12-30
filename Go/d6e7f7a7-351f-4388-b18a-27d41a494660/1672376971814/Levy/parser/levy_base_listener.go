// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376971814/Levy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Levy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLevyListener is a complete listener for a parse tree produced by LevyParser.
type BaseLevyListener struct{}

var _ LevyListener = &BaseLevyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLevyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLevyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLevyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLevyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLevyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLevyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLevy is called when production levy is entered.
func (s *BaseLevyListener) EnterLevy(ctx *LevyContext) {}

// ExitLevy is called when production levy is exited.
func (s *BaseLevyListener) ExitLevy(ctx *LevyContext) {}
