// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669621222629/Flag.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Flag

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFlagListener is a complete listener for a parse tree produced by FlagParser.
type BaseFlagListener struct{}

var _ FlagListener = &BaseFlagListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFlagListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFlagListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFlagListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFlagListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFlagListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFlagListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFlag is called when production flag is entered.
func (s *BaseFlagListener) EnterFlag(ctx *FlagContext) {}

// ExitFlag is called when production flag is exited.
func (s *BaseFlagListener) ExitFlag(ctx *FlagContext) {}
