// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669793549359/ShapeFactor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ShapeFactor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseShapeFactorListener is a complete listener for a parse tree produced by ShapeFactorParser.
type BaseShapeFactorListener struct{}

var _ ShapeFactorListener = &BaseShapeFactorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseShapeFactorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseShapeFactorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseShapeFactorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseShapeFactorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseShapeFactorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseShapeFactorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterShapefactor is called when production shapefactor is entered.
func (s *BaseShapeFactorListener) EnterShapefactor(ctx *ShapefactorContext) {}

// ExitShapefactor is called when production shapefactor is exited.
func (s *BaseShapeFactorListener) ExitShapefactor(ctx *ShapefactorContext) {}
