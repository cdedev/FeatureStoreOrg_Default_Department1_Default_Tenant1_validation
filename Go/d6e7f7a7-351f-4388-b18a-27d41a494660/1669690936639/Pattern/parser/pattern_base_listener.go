// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690936639/Pattern.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pattern

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePatternListener is a complete listener for a parse tree produced by PatternParser.
type BasePatternListener struct{}

var _ PatternListener = &BasePatternListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePatternListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePatternListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePatternListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePatternListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePatternListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePatternListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BasePatternListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BasePatternListener) ExitPattern(ctx *PatternContext) {}
