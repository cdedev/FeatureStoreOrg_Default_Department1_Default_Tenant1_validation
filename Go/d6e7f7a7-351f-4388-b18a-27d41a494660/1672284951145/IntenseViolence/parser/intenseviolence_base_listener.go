// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284951145/IntenseViolence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IntenseViolence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIntenseViolenceListener is a complete listener for a parse tree produced by IntenseViolenceParser.
type BaseIntenseViolenceListener struct{}

var _ IntenseViolenceListener = &BaseIntenseViolenceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIntenseViolenceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIntenseViolenceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIntenseViolenceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIntenseViolenceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIntenseViolenceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIntenseViolenceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIntenseviolence is called when production intenseviolence is entered.
func (s *BaseIntenseViolenceListener) EnterIntenseviolence(ctx *IntenseviolenceContext) {}

// ExitIntenseviolence is called when production intenseviolence is exited.
func (s *BaseIntenseViolenceListener) ExitIntenseviolence(ctx *IntenseviolenceContext) {}
