// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673676034823/Dimensions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dimensions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDimensionsListener is a complete listener for a parse tree produced by DimensionsParser.
type BaseDimensionsListener struct{}

var _ DimensionsListener = &BaseDimensionsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDimensionsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDimensionsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDimensionsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDimensionsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDimensionsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDimensionsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDimensions is called when production dimensions is entered.
func (s *BaseDimensionsListener) EnterDimensions(ctx *DimensionsContext) {}

// ExitDimensions is called when production dimensions is exited.
func (s *BaseDimensionsListener) ExitDimensions(ctx *DimensionsContext) {}
