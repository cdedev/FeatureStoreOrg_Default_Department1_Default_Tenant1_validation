// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671599159281/Entity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Entity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEntityListener is a complete listener for a parse tree produced by EntityParser.
type BaseEntityListener struct{}

var _ EntityListener = &BaseEntityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEntityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEntityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEntityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEntityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEntityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEntityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEntity is called when production entity is entered.
func (s *BaseEntityListener) EnterEntity(ctx *EntityContext) {}

// ExitEntity is called when production entity is exited.
func (s *BaseEntityListener) ExitEntity(ctx *EntityContext) {}
