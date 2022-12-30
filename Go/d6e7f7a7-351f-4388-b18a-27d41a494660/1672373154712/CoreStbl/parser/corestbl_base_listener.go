// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672373154712/CoreStbl.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CoreStbl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCoreStblListener is a complete listener for a parse tree produced by CoreStblParser.
type BaseCoreStblListener struct{}

var _ CoreStblListener = &BaseCoreStblListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCoreStblListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCoreStblListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCoreStblListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCoreStblListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCoreStblListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCoreStblListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCorestbl is called when production corestbl is entered.
func (s *BaseCoreStblListener) EnterCorestbl(ctx *CorestblContext) {}

// ExitCorestbl is called when production corestbl is exited.
func (s *BaseCoreStblListener) ExitCorestbl(ctx *CorestblContext) {}
