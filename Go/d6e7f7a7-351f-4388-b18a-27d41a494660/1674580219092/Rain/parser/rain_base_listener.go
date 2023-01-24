// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674580219092/Rain.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rain

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRainListener is a complete listener for a parse tree produced by RainParser.
type BaseRainListener struct{}

var _ RainListener = &BaseRainListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRainListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRainListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRainListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRainListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRainListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRainListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRain is called when production rain is entered.
func (s *BaseRainListener) EnterRain(ctx *RainContext) {}

// ExitRain is called when production rain is exited.
func (s *BaseRainListener) ExitRain(ctx *RainContext) {}
