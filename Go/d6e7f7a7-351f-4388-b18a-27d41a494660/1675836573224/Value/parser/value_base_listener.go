// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675836573224/Value.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Value

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseValueListener is a complete listener for a parse tree produced by ValueParser.
type BaseValueListener struct{}

var _ ValueListener = &BaseValueListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseValueListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseValueListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseValueListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseValueListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseValueListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseValueListener) ExitExpression(ctx *ExpressionContext) {}

// EnterValue is called when production value is entered.
func (s *BaseValueListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseValueListener) ExitValue(ctx *ValueContext) {}
