// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676609525407/Quality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Quality

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQualityListener is a complete listener for a parse tree produced by QualityParser.
type BaseQualityListener struct{}

var _ QualityListener = &BaseQualityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQualityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQualityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQualityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQualityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseQualityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseQualityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterQuality is called when production quality is entered.
func (s *BaseQualityListener) EnterQuality(ctx *QualityContext) {}

// ExitQuality is called when production quality is exited.
func (s *BaseQualityListener) ExitQuality(ctx *QualityContext) {}
