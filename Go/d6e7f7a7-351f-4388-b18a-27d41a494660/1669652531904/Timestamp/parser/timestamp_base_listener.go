// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669652531904/Timestamp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Timestamp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTimestampListener is a complete listener for a parse tree produced by TimestampParser.
type BaseTimestampListener struct{}

var _ TimestampListener = &BaseTimestampListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTimestampListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTimestampListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTimestampListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTimestampListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTimestampListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTimestampListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTimestamp is called when production timestamp is entered.
func (s *BaseTimestampListener) EnterTimestamp(ctx *TimestampContext) {}

// ExitTimestamp is called when production timestamp is exited.
func (s *BaseTimestampListener) ExitTimestamp(ctx *TimestampContext) {}
