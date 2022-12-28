// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205472033/PoojaRoom.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PoojaRoom

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePoojaRoomListener is a complete listener for a parse tree produced by PoojaRoomParser.
type BasePoojaRoomListener struct{}

var _ PoojaRoomListener = &BasePoojaRoomListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePoojaRoomListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePoojaRoomListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePoojaRoomListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePoojaRoomListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePoojaRoomListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePoojaRoomListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPoojaroom is called when production poojaroom is entered.
func (s *BasePoojaRoomListener) EnterPoojaroom(ctx *PoojaroomContext) {}

// ExitPoojaroom is called when production poojaroom is exited.
func (s *BasePoojaRoomListener) ExitPoojaroom(ctx *PoojaroomContext) {}
