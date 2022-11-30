// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669792843997/Width.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Width

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWidthListener is a complete listener for a parse tree produced by WidthParser.
type BaseWidthListener struct{}

var _ WidthListener = &BaseWidthListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWidthListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWidthListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWidthListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWidthListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWidthListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWidthListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWidth is called when production width is entered.
func (s *BaseWidthListener) EnterWidth(ctx *WidthContext) {}

// ExitWidth is called when production width is exited.
func (s *BaseWidthListener) ExitWidth(ctx *WidthContext) {}
