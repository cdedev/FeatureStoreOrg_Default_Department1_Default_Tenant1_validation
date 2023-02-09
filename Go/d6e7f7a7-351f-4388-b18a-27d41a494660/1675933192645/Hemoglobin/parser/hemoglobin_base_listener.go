// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675933192645/Hemoglobin.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hemoglobin

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHemoglobinListener is a complete listener for a parse tree produced by HemoglobinParser.
type BaseHemoglobinListener struct{}

var _ HemoglobinListener = &BaseHemoglobinListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHemoglobinListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHemoglobinListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHemoglobinListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHemoglobinListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHemoglobinListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHemoglobinListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHemoglobin is called when production hemoglobin is entered.
func (s *BaseHemoglobinListener) EnterHemoglobin(ctx *HemoglobinContext) {}

// ExitHemoglobin is called when production hemoglobin is exited.
func (s *BaseHemoglobinListener) ExitHemoglobin(ctx *HemoglobinContext) {}
