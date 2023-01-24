// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674546402043/StatusOne.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // StatusOne

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStatusOneListener is a complete listener for a parse tree produced by StatusOneParser.
type BaseStatusOneListener struct{}

var _ StatusOneListener = &BaseStatusOneListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStatusOneListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStatusOneListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStatusOneListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStatusOneListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStatusOneListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStatusOneListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStatusone is called when production statusone is entered.
func (s *BaseStatusOneListener) EnterStatusone(ctx *StatusoneContext) {}

// ExitStatusone is called when production statusone is exited.
func (s *BaseStatusOneListener) ExitStatusone(ctx *StatusoneContext) {}
