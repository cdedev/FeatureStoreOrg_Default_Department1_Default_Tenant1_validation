// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676515923990/Injured.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Injured

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInjuredListener is a complete listener for a parse tree produced by InjuredParser.
type BaseInjuredListener struct{}

var _ InjuredListener = &BaseInjuredListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInjuredListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInjuredListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInjuredListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInjuredListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInjuredListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInjuredListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInjured is called when production injured is entered.
func (s *BaseInjuredListener) EnterInjured(ctx *InjuredContext) {}

// ExitInjured is called when production injured is exited.
func (s *BaseInjuredListener) ExitInjured(ctx *InjuredContext) {}
