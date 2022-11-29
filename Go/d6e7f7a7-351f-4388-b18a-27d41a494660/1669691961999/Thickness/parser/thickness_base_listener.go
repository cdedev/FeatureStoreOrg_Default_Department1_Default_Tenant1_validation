// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669691961999/Thickness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Thickness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseThicknessListener is a complete listener for a parse tree produced by ThicknessParser.
type BaseThicknessListener struct{}

var _ ThicknessListener = &BaseThicknessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseThicknessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseThicknessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseThicknessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseThicknessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseThicknessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseThicknessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterThickness is called when production thickness is entered.
func (s *BaseThicknessListener) EnterThickness(ctx *ThicknessContext) {}

// ExitThickness is called when production thickness is exited.
func (s *BaseThicknessListener) ExitThickness(ctx *ThicknessContext) {}
