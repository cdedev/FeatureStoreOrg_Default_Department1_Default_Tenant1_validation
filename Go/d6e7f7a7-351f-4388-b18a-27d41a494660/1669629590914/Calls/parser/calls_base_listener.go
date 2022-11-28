// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669629590914/Calls.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Calls

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCallsListener is a complete listener for a parse tree produced by CallsParser.
type BaseCallsListener struct{}

var _ CallsListener = &BaseCallsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCallsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCallsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCallsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCallsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCallsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCallsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCalls is called when production calls is entered.
func (s *BaseCallsListener) EnterCalls(ctx *CallsContext) {}

// ExitCalls is called when production calls is exited.
func (s *BaseCallsListener) ExitCalls(ctx *CallsContext) {}
