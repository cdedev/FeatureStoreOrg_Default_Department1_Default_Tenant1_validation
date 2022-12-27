// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118879751/MaxSpeed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MaxSpeed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMaxSpeedListener is a complete listener for a parse tree produced by MaxSpeedParser.
type BaseMaxSpeedListener struct{}

var _ MaxSpeedListener = &BaseMaxSpeedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMaxSpeedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMaxSpeedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMaxSpeedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMaxSpeedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMaxSpeedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMaxSpeedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMaxspeed is called when production maxspeed is entered.
func (s *BaseMaxSpeedListener) EnterMaxspeed(ctx *MaxspeedContext) {}

// ExitMaxspeed is called when production maxspeed is exited.
func (s *BaseMaxSpeedListener) ExitMaxspeed(ctx *MaxspeedContext) {}
