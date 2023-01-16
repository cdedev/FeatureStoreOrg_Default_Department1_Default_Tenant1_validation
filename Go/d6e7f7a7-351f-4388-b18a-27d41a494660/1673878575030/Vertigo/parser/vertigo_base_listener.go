// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878575030/Vertigo.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vertigo

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVertigoListener is a complete listener for a parse tree produced by VertigoParser.
type BaseVertigoListener struct{}

var _ VertigoListener = &BaseVertigoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVertigoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVertigoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVertigoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVertigoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVertigoListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVertigoListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVertigo is called when production vertigo is entered.
func (s *BaseVertigoListener) EnterVertigo(ctx *VertigoContext) {}

// ExitVertigo is called when production vertigo is exited.
func (s *BaseVertigoListener) ExitVertigo(ctx *VertigoContext) {}
