// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205383372/StudyRoom.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // StudyRoom

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStudyRoomListener is a complete listener for a parse tree produced by StudyRoomParser.
type BaseStudyRoomListener struct{}

var _ StudyRoomListener = &BaseStudyRoomListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStudyRoomListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStudyRoomListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStudyRoomListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStudyRoomListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStudyRoomListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStudyRoomListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStudyroom is called when production studyroom is entered.
func (s *BaseStudyRoomListener) EnterStudyroom(ctx *StudyroomContext) {}

// ExitStudyroom is called when production studyroom is exited.
func (s *BaseStudyRoomListener) ExitStudyroom(ctx *StudyroomContext) {}
