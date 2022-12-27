// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118775991/Nausea.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nausea

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNauseaListener is a complete listener for a parse tree produced by NauseaParser.
type BaseNauseaListener struct{}

var _ NauseaListener = &BaseNauseaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNauseaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNauseaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNauseaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNauseaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNauseaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNauseaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNausea is called when production nausea is entered.
func (s *BaseNauseaListener) EnterNausea(ctx *NauseaContext) {}

// ExitNausea is called when production nausea is exited.
func (s *BaseNauseaListener) ExitNausea(ctx *NauseaContext) {}
