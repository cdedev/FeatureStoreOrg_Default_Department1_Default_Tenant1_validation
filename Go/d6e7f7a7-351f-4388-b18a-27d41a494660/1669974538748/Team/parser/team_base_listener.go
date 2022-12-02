// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669974538748/Team.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Team

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTeamListener is a complete listener for a parse tree produced by TeamParser.
type BaseTeamListener struct{}

var _ TeamListener = &BaseTeamListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTeamListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTeamListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTeamListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTeamListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTeamListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTeamListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTeam is called when production team is entered.
func (s *BaseTeamListener) EnterTeam(ctx *TeamContext) {}

// ExitTeam is called when production team is exited.
func (s *BaseTeamListener) ExitTeam(ctx *TeamContext) {}
