// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672115355714/Email.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Email

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEmailListener is a complete listener for a parse tree produced by EmailParser.
type BaseEmailListener struct{}

var _ EmailListener = &BaseEmailListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEmailListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEmailListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEmailListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEmailListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEmailListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEmailListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEmail is called when production email is entered.
func (s *BaseEmailListener) EnterEmail(ctx *EmailContext) {}

// ExitEmail is called when production email is exited.
func (s *BaseEmailListener) ExitEmail(ctx *EmailContext) {}
