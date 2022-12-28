// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205428042/ServantRoom.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ServantRoom

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseServantRoomListener is a complete listener for a parse tree produced by ServantRoomParser.
type BaseServantRoomListener struct{}

var _ ServantRoomListener = &BaseServantRoomListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseServantRoomListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseServantRoomListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseServantRoomListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseServantRoomListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseServantRoomListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseServantRoomListener) ExitExpression(ctx *ExpressionContext) {}

// EnterServantroom is called when production servantroom is entered.
func (s *BaseServantRoomListener) EnterServantroom(ctx *ServantroomContext) {}

// ExitServantroom is called when production servantroom is exited.
func (s *BaseServantRoomListener) ExitServantroom(ctx *ServantroomContext) {}
