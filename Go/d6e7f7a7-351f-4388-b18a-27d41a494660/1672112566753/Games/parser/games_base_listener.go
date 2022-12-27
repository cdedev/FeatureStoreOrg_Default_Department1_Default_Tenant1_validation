// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672112566753/Games.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Games

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGamesListener is a complete listener for a parse tree produced by GamesParser.
type BaseGamesListener struct{}

var _ GamesListener = &BaseGamesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGamesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGamesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGamesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGamesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGamesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGamesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGames is called when production games is entered.
func (s *BaseGamesListener) EnterGames(ctx *GamesContext) {}

// ExitGames is called when production games is exited.
func (s *BaseGamesListener) ExitGames(ctx *GamesContext) {}
