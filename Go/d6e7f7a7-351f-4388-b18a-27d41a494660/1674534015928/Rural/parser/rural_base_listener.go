// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674534015928/Rural.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rural

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRuralListener is a complete listener for a parse tree produced by RuralParser.
type BaseRuralListener struct{}

var _ RuralListener = &BaseRuralListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRuralListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRuralListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRuralListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRuralListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRuralListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRuralListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRural is called when production rural is entered.
func (s *BaseRuralListener) EnterRural(ctx *RuralContext) {}

// ExitRural is called when production rural is exited.
func (s *BaseRuralListener) ExitRural(ctx *RuralContext) {}
