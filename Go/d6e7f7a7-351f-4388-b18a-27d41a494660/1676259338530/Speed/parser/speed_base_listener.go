// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676259338530/Speed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Speed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSpeedListener is a complete listener for a parse tree produced by SpeedParser.
type BaseSpeedListener struct{}

var _ SpeedListener = &BaseSpeedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSpeedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSpeedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSpeedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSpeedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSpeedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSpeedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSpeed is called when production speed is entered.
func (s *BaseSpeedListener) EnterSpeed(ctx *SpeedContext) {}

// ExitSpeed is called when production speed is exited.
func (s *BaseSpeedListener) ExitSpeed(ctx *SpeedContext) {}
