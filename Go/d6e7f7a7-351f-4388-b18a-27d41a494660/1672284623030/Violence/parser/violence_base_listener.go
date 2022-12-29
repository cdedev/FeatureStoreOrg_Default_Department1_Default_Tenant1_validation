// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284623030/Violence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Violence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseViolenceListener is a complete listener for a parse tree produced by ViolenceParser.
type BaseViolenceListener struct{}

var _ ViolenceListener = &BaseViolenceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseViolenceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseViolenceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseViolenceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseViolenceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseViolenceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseViolenceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterViolence is called when production violence is entered.
func (s *BaseViolenceListener) EnterViolence(ctx *ViolenceContext) {}

// ExitViolence is called when production violence is exited.
func (s *BaseViolenceListener) ExitViolence(ctx *ViolenceContext) {}
