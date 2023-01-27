// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674792977494/Community.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Community

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCommunityListener is a complete listener for a parse tree produced by CommunityParser.
type BaseCommunityListener struct{}

var _ CommunityListener = &BaseCommunityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCommunityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCommunityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCommunityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCommunityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCommunityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCommunityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCommunity is called when production community is entered.
func (s *BaseCommunityListener) EnterCommunity(ctx *CommunityContext) {}

// ExitCommunity is called when production community is exited.
func (s *BaseCommunityListener) ExitCommunity(ctx *CommunityContext) {}
