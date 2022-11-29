// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669693636336/Room.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Room

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRoomListener is a complete listener for a parse tree produced by RoomParser.
type BaseRoomListener struct{}

var _ RoomListener = &BaseRoomListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRoomListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRoomListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRoomListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRoomListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRoomListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRoomListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRoom is called when production room is entered.
func (s *BaseRoomListener) EnterRoom(ctx *RoomContext) {}

// ExitRoom is called when production room is exited.
func (s *BaseRoomListener) ExitRoom(ctx *RoomContext) {}
