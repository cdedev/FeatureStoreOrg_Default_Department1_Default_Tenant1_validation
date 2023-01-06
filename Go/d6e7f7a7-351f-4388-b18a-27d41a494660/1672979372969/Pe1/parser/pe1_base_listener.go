// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672979372969/Pe1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pe1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePe1Listener is a complete listener for a parse tree produced by Pe1Parser.
type BasePe1Listener struct{}

var _ Pe1Listener = &BasePe1Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePe1Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePe1Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePe1Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePe1Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePe1Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePe1Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterPe1 is called when production pe1 is entered.
func (s *BasePe1Listener) EnterPe1(ctx *Pe1Context) {}

// ExitPe1 is called when production pe1 is exited.
func (s *BasePe1Listener) ExitPe1(ctx *Pe1Context) {}
