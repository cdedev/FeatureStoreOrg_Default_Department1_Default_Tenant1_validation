// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672979442491/Tp1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tp1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTp1Listener is a complete listener for a parse tree produced by Tp1Parser.
type BaseTp1Listener struct{}

var _ Tp1Listener = &BaseTp1Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTp1Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTp1Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTp1Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTp1Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTp1Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTp1Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterTp1 is called when production tp1 is entered.
func (s *BaseTp1Listener) EnterTp1(ctx *Tp1Context) {}

// ExitTp1 is called when production tp1 is exited.
func (s *BaseTp1Listener) ExitTp1(ctx *Tp1Context) {}
