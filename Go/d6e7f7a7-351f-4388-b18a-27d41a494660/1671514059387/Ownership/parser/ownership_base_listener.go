// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671514059387/Ownership.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ownership

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOwnershipListener is a complete listener for a parse tree produced by OwnershipParser.
type BaseOwnershipListener struct{}

var _ OwnershipListener = &BaseOwnershipListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOwnershipListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOwnershipListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOwnershipListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOwnershipListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOwnershipListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOwnershipListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOwnership is called when production ownership is entered.
func (s *BaseOwnershipListener) EnterOwnership(ctx *OwnershipContext) {}

// ExitOwnership is called when production ownership is exited.
func (s *BaseOwnershipListener) ExitOwnership(ctx *OwnershipContext) {}
