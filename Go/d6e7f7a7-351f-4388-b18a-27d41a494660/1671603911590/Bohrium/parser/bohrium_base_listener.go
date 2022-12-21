// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603911590/Bohrium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bohrium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBohriumListener is a complete listener for a parse tree produced by BohriumParser.
type BaseBohriumListener struct{}

var _ BohriumListener = &BaseBohriumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBohriumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBohriumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBohriumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBohriumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBohriumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBohriumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBohrium is called when production bohrium is entered.
func (s *BaseBohriumListener) EnterBohrium(ctx *BohriumContext) {}

// ExitBohrium is called when production bohrium is exited.
func (s *BaseBohriumListener) ExitBohrium(ctx *BohriumContext) {}
