// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675830054573/Followers.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Followers

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFollowersListener is a complete listener for a parse tree produced by FollowersParser.
type BaseFollowersListener struct{}

var _ FollowersListener = &BaseFollowersListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFollowersListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFollowersListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFollowersListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFollowersListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFollowersListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFollowersListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFollowers is called when production followers is entered.
func (s *BaseFollowersListener) EnterFollowers(ctx *FollowersContext) {}

// ExitFollowers is called when production followers is exited.
func (s *BaseFollowersListener) ExitFollowers(ctx *FollowersContext) {}
