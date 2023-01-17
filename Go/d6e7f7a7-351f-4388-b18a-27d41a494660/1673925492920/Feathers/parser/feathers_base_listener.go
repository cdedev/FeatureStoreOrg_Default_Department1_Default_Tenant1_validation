// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925492920/Feathers.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Feathers

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFeathersListener is a complete listener for a parse tree produced by FeathersParser.
type BaseFeathersListener struct{}

var _ FeathersListener = &BaseFeathersListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFeathersListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFeathersListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFeathersListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFeathersListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFeathersListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFeathersListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFeathers is called when production feathers is entered.
func (s *BaseFeathersListener) EnterFeathers(ctx *FeathersContext) {}

// ExitFeathers is called when production feathers is exited.
func (s *BaseFeathersListener) ExitFeathers(ctx *FeathersContext) {}
