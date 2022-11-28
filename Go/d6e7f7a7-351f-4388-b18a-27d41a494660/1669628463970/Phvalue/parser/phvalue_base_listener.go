// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628463970/Phvalue.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Phvalue

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePhvalueListener is a complete listener for a parse tree produced by PhvalueParser.
type BasePhvalueListener struct{}

var _ PhvalueListener = &BasePhvalueListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePhvalueListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePhvalueListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePhvalueListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePhvalueListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePhvalueListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePhvalueListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPhvalue is called when production phvalue is entered.
func (s *BasePhvalueListener) EnterPhvalue(ctx *PhvalueContext) {}

// ExitPhvalue is called when production phvalue is exited.
func (s *BasePhvalueListener) ExitPhvalue(ctx *PhvalueContext) {}
