// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675936356062/Lifespan.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lifespan

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLifespanListener is a complete listener for a parse tree produced by LifespanParser.
type BaseLifespanListener struct{}

var _ LifespanListener = &BaseLifespanListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLifespanListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLifespanListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLifespanListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLifespanListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLifespanListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLifespanListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLifespan is called when production lifespan is entered.
func (s *BaseLifespanListener) EnterLifespan(ctx *LifespanContext) {}

// ExitLifespan is called when production lifespan is exited.
func (s *BaseLifespanListener) ExitLifespan(ctx *LifespanContext) {}
