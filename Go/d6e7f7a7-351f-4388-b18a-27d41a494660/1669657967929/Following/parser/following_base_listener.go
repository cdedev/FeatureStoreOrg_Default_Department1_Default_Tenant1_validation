// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669657967929/Following.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Following

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFollowingListener is a complete listener for a parse tree produced by FollowingParser.
type BaseFollowingListener struct{}

var _ FollowingListener = &BaseFollowingListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFollowingListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFollowingListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFollowingListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFollowingListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFollowingListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFollowingListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFollowing is called when production following is entered.
func (s *BaseFollowingListener) EnterFollowing(ctx *FollowingContext) {}

// ExitFollowing is called when production following is exited.
func (s *BaseFollowingListener) ExitFollowing(ctx *FollowingContext) {}
