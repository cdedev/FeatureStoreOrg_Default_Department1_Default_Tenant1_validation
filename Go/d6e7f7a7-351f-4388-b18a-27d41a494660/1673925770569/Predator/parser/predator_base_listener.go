// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925770569/Predator.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Predator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePredatorListener is a complete listener for a parse tree produced by PredatorParser.
type BasePredatorListener struct{}

var _ PredatorListener = &BasePredatorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePredatorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePredatorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePredatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePredatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePredatorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePredatorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPredator is called when production predator is entered.
func (s *BasePredatorListener) EnterPredator(ctx *PredatorContext) {}

// ExitPredator is called when production predator is exited.
func (s *BasePredatorListener) ExitPredator(ctx *PredatorContext) {}
