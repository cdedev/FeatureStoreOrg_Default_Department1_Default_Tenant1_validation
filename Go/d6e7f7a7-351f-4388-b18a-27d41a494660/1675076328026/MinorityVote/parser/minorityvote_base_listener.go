// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675076328026/MinorityVote.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MinorityVote

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMinorityVoteListener is a complete listener for a parse tree produced by MinorityVoteParser.
type BaseMinorityVoteListener struct{}

var _ MinorityVoteListener = &BaseMinorityVoteListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMinorityVoteListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMinorityVoteListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMinorityVoteListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMinorityVoteListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMinorityVoteListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMinorityVoteListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMinorityvote is called when production minorityvote is entered.
func (s *BaseMinorityVoteListener) EnterMinorityvote(ctx *MinorityvoteContext) {}

// ExitMinorityvote is called when production minorityvote is exited.
func (s *BaseMinorityVoteListener) ExitMinorityvote(ctx *MinorityvoteContext) {}
