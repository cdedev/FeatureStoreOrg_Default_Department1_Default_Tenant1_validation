// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675834897679/Posture.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Posture

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePostureListener is a complete listener for a parse tree produced by PostureParser.
type BasePostureListener struct{}

var _ PostureListener = &BasePostureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePostureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePostureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePostureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePostureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePostureListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePostureListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPosture is called when production posture is entered.
func (s *BasePostureListener) EnterPosture(ctx *PostureContext) {}

// ExitPosture is called when production posture is exited.
func (s *BasePostureListener) ExitPosture(ctx *PostureContext) {}
