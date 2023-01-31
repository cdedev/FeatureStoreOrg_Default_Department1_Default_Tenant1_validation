// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675145759639/Shape.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Shape

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseShapeListener is a complete listener for a parse tree produced by ShapeParser.
type BaseShapeListener struct{}

var _ ShapeListener = &BaseShapeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseShapeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseShapeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseShapeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseShapeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterShape is called when production shape is entered.
func (s *BaseShapeListener) EnterShape(ctx *ShapeContext) {}

// ExitShape is called when production shape is exited.
func (s *BaseShapeListener) ExitShape(ctx *ShapeContext) {}
