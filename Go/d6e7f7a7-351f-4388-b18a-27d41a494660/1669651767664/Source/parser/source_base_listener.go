// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669651767664/Source.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Source

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSourceListener is a complete listener for a parse tree produced by SourceParser.
type BaseSourceListener struct{}

var _ SourceListener = &BaseSourceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSourceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSourceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSourceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSourceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSourceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSourceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSource is called when production source is entered.
func (s *BaseSourceListener) EnterSource(ctx *SourceContext) {}

// ExitSource is called when production source is exited.
func (s *BaseSourceListener) ExitSource(ctx *SourceContext) {}
