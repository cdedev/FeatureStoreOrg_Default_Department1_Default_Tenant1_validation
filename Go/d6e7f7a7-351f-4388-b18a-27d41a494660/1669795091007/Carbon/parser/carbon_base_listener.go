// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669795091007/Carbon.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carbon

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCarbonListener is a complete listener for a parse tree produced by CarbonParser.
type BaseCarbonListener struct{}

var _ CarbonListener = &BaseCarbonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCarbonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCarbonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCarbonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCarbonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCarbonListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCarbonListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCarbon is called when production carbon is entered.
func (s *BaseCarbonListener) EnterCarbon(ctx *CarbonContext) {}

// ExitCarbon is called when production carbon is exited.
func (s *BaseCarbonListener) ExitCarbon(ctx *CarbonContext) {}
