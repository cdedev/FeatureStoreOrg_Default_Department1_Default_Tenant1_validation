// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673675990685/Concavity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Concavity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConcavityListener is a complete listener for a parse tree produced by ConcavityParser.
type BaseConcavityListener struct{}

var _ ConcavityListener = &BaseConcavityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConcavityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConcavityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConcavityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConcavityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConcavityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConcavityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConcavity is called when production concavity is entered.
func (s *BaseConcavityListener) EnterConcavity(ctx *ConcavityContext) {}

// ExitConcavity is called when production concavity is exited.
func (s *BaseConcavityListener) ExitConcavity(ctx *ConcavityContext) {}
