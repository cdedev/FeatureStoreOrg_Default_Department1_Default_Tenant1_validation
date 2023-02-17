// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676608227546/Satisfaction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Satisfaction

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSatisfactionListener is a complete listener for a parse tree produced by SatisfactionParser.
type BaseSatisfactionListener struct{}

var _ SatisfactionListener = &BaseSatisfactionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSatisfactionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSatisfactionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSatisfactionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSatisfactionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSatisfactionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSatisfactionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSatisfaction is called when production satisfaction is entered.
func (s *BaseSatisfactionListener) EnterSatisfaction(ctx *SatisfactionContext) {}

// ExitSatisfaction is called when production satisfaction is exited.
func (s *BaseSatisfactionListener) ExitSatisfaction(ctx *SatisfactionContext) {}
