// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604495022/Cobalt.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cobalt

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCobaltListener is a complete listener for a parse tree produced by CobaltParser.
type BaseCobaltListener struct{}

var _ CobaltListener = &BaseCobaltListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCobaltListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCobaltListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCobaltListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCobaltListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCobaltListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCobaltListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCobalt is called when production cobalt is entered.
func (s *BaseCobaltListener) EnterCobalt(ctx *CobaltContext) {}

// ExitCobalt is called when production cobalt is exited.
func (s *BaseCobaltListener) ExitCobalt(ctx *CobaltContext) {}
