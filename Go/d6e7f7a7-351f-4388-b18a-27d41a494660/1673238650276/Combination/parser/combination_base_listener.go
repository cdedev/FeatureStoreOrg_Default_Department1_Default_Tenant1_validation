// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238650276/Combination.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Combination

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCombinationListener is a complete listener for a parse tree produced by CombinationParser.
type BaseCombinationListener struct{}

var _ CombinationListener = &BaseCombinationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCombinationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCombinationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCombinationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCombinationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCombinationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCombinationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCombination is called when production combination is entered.
func (s *BaseCombinationListener) EnterCombination(ctx *CombinationContext) {}

// ExitCombination is called when production combination is exited.
func (s *BaseCombinationListener) ExitCombination(ctx *CombinationContext) {}
