// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972122105/Close.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Close

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCloseListener is a complete listener for a parse tree produced by CloseParser.
type BaseCloseListener struct{}

var _ CloseListener = &BaseCloseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCloseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCloseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCloseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCloseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCloseListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCloseListener) ExitExpression(ctx *ExpressionContext) {}

// EnterClose is called when production close is entered.
func (s *BaseCloseListener) EnterClose(ctx *CloseContext) {}

// ExitClose is called when production close is exited.
func (s *BaseCloseListener) ExitClose(ctx *CloseContext) {}
