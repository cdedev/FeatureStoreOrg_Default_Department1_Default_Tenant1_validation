// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669788081665/Direction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Direction

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDirectionListener is a complete listener for a parse tree produced by DirectionParser.
type BaseDirectionListener struct{}

var _ DirectionListener = &BaseDirectionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDirectionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDirectionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDirectionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDirectionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDirectionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDirectionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDirection is called when production direction is entered.
func (s *BaseDirectionListener) EnterDirection(ctx *DirectionContext) {}

// ExitDirection is called when production direction is exited.
func (s *BaseDirectionListener) ExitDirection(ctx *DirectionContext) {}
