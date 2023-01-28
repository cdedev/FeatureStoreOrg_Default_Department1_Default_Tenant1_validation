// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674881836211/Frozen.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Frozen

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFrozenListener is a complete listener for a parse tree produced by FrozenParser.
type BaseFrozenListener struct{}

var _ FrozenListener = &BaseFrozenListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFrozenListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFrozenListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFrozenListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFrozenListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFrozenListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFrozenListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFrozen is called when production frozen is entered.
func (s *BaseFrozenListener) EnterFrozen(ctx *FrozenContext) {}

// ExitFrozen is called when production frozen is exited.
func (s *BaseFrozenListener) ExitFrozen(ctx *FrozenContext) {}
