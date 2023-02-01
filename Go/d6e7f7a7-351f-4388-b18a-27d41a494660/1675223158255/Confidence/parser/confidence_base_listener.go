// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675223158255/Confidence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Confidence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConfidenceListener is a complete listener for a parse tree produced by ConfidenceParser.
type BaseConfidenceListener struct{}

var _ ConfidenceListener = &BaseConfidenceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConfidenceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConfidenceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConfidenceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConfidenceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConfidenceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConfidenceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConfidence is called when production confidence is entered.
func (s *BaseConfidenceListener) EnterConfidence(ctx *ConfidenceContext) {}

// ExitConfidence is called when production confidence is exited.
func (s *BaseConfidenceListener) ExitConfidence(ctx *ConfidenceContext) {}
