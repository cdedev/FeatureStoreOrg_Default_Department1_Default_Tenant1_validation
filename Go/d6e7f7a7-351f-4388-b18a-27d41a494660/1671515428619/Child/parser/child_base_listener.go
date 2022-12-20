// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671515428619/Child.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Child

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChildListener is a complete listener for a parse tree produced by ChildParser.
type BaseChildListener struct{}

var _ ChildListener = &BaseChildListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChildListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChildListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChildListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChildListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseChildListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseChildListener) ExitExpression(ctx *ExpressionContext) {}

// EnterChild is called when production child is entered.
func (s *BaseChildListener) EnterChild(ctx *ChildContext) {}

// ExitChild is called when production child is exited.
func (s *BaseChildListener) ExitChild(ctx *ChildContext) {}
