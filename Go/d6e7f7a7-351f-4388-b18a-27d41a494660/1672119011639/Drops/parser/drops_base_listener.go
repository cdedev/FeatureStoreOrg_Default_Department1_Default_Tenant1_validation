// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119011639/Drops.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Drops

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDropsListener is a complete listener for a parse tree produced by DropsParser.
type BaseDropsListener struct{}

var _ DropsListener = &BaseDropsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDropsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDropsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDropsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDropsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDropsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDropsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDrops is called when production drops is entered.
func (s *BaseDropsListener) EnterDrops(ctx *DropsContext) {}

// ExitDrops is called when production drops is exited.
func (s *BaseDropsListener) ExitDrops(ctx *DropsContext) {}
