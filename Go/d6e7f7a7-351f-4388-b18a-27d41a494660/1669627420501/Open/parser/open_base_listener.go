// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669627420501/Open.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Open

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOpenListener is a complete listener for a parse tree produced by OpenParser.
type BaseOpenListener struct{}

var _ OpenListener = &BaseOpenListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOpenListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOpenListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOpenListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOpenListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOpenListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOpenListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOpen is called when production open is entered.
func (s *BaseOpenListener) EnterOpen(ctx *OpenContext) {}

// ExitOpen is called when production open is exited.
func (s *BaseOpenListener) ExitOpen(ctx *OpenContext) {}
