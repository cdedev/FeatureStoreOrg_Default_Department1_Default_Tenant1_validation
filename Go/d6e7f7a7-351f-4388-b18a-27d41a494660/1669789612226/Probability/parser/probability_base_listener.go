// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669789612226/Probability.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Probability

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProbabilityListener is a complete listener for a parse tree produced by ProbabilityParser.
type BaseProbabilityListener struct{}

var _ ProbabilityListener = &BaseProbabilityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProbabilityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProbabilityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProbabilityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProbabilityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseProbabilityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseProbabilityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterProbability is called when production probability is entered.
func (s *BaseProbabilityListener) EnterProbability(ctx *ProbabilityContext) {}

// ExitProbability is called when production probability is exited.
func (s *BaseProbabilityListener) ExitProbability(ctx *ProbabilityContext) {}
