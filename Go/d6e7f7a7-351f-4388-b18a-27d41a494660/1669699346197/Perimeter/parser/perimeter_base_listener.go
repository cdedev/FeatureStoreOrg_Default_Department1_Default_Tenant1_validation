// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669699346197/Perimeter.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Perimeter

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePerimeterListener is a complete listener for a parse tree produced by PerimeterParser.
type BasePerimeterListener struct{}

var _ PerimeterListener = &BasePerimeterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePerimeterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePerimeterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePerimeterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePerimeterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePerimeterListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePerimeterListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPerimeter is called when production perimeter is entered.
func (s *BasePerimeterListener) EnterPerimeter(ctx *PerimeterContext) {}

// ExitPerimeter is called when production perimeter is exited.
func (s *BasePerimeterListener) ExitPerimeter(ctx *PerimeterContext) {}
