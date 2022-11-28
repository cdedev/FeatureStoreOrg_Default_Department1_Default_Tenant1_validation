// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624525960/Measure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Measure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMeasureListener is a complete listener for a parse tree produced by MeasureParser.
type BaseMeasureListener struct{}

var _ MeasureListener = &BaseMeasureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMeasureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMeasureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMeasureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMeasureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMeasureListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMeasureListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMeasure is called when production measure is entered.
func (s *BaseMeasureListener) EnterMeasure(ctx *MeasureContext) {}

// ExitMeasure is called when production measure is exited.
func (s *BaseMeasureListener) ExitMeasure(ctx *MeasureContext) {}
