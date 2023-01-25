// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674666543500/Gamma.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gamma

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGammaListener is a complete listener for a parse tree produced by GammaParser.
type BaseGammaListener struct{}

var _ GammaListener = &BaseGammaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGammaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGammaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGammaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGammaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGammaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGammaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGamma is called when production gamma is entered.
func (s *BaseGammaListener) EnterGamma(ctx *GammaContext) {}

// ExitGamma is called when production gamma is exited.
func (s *BaseGammaListener) ExitGamma(ctx *GammaContext) {}
