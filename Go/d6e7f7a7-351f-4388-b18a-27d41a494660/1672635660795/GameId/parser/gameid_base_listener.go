// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672635660795/GameId.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GameId

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGameIdListener is a complete listener for a parse tree produced by GameIdParser.
type BaseGameIdListener struct{}

var _ GameIdListener = &BaseGameIdListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGameIdListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGameIdListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGameIdListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGameIdListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGameIdListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGameIdListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGameid is called when production gameid is entered.
func (s *BaseGameIdListener) EnterGameid(ctx *GameidContext) {}

// ExitGameid is called when production gameid is exited.
func (s *BaseGameIdListener) ExitGameid(ctx *GameidContext) {}
