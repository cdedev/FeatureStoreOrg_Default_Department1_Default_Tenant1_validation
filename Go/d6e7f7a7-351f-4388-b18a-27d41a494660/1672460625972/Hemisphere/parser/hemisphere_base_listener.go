// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672460625972/Hemisphere.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hemisphere

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHemisphereListener is a complete listener for a parse tree produced by HemisphereParser.
type BaseHemisphereListener struct{}

var _ HemisphereListener = &BaseHemisphereListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHemisphereListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHemisphereListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHemisphereListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHemisphereListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHemisphereListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHemisphereListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHemisphere is called when production hemisphere is entered.
func (s *BaseHemisphereListener) EnterHemisphere(ctx *HemisphereContext) {}

// ExitHemisphere is called when production hemisphere is exited.
func (s *BaseHemisphereListener) ExitHemisphere(ctx *HemisphereContext) {}
