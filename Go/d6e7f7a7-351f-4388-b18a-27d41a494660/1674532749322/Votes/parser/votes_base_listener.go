// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674532749322/Votes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Votes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVotesListener is a complete listener for a parse tree produced by VotesParser.
type BaseVotesListener struct{}

var _ VotesListener = &BaseVotesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVotesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVotesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVotesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVotesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVotesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVotesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVotes is called when production votes is entered.
func (s *BaseVotesListener) EnterVotes(ctx *VotesContext) {}

// ExitVotes is called when production votes is exited.
func (s *BaseVotesListener) ExitVotes(ctx *VotesContext) {}
