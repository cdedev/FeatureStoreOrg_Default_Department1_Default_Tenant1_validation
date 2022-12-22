// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671697563529/Children.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Children

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChildrenListener is a complete listener for a parse tree produced by ChildrenParser.
type BaseChildrenListener struct{}

var _ ChildrenListener = &BaseChildrenListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChildrenListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChildrenListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChildrenListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChildrenListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseChildrenListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseChildrenListener) ExitExpression(ctx *ExpressionContext) {}

// EnterChildren is called when production children is entered.
func (s *BaseChildrenListener) EnterChildren(ctx *ChildrenContext) {}

// ExitChildren is called when production children is exited.
func (s *BaseChildrenListener) ExitChildren(ctx *ChildrenContext) {}
