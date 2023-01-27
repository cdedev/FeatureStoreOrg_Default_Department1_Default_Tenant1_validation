// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674798267059/Turnover.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Turnover

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTurnoverListener is a complete listener for a parse tree produced by TurnoverParser.
type BaseTurnoverListener struct{}

var _ TurnoverListener = &BaseTurnoverListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTurnoverListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTurnoverListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTurnoverListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTurnoverListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTurnoverListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTurnoverListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTurnover is called when production turnover is entered.
func (s *BaseTurnoverListener) EnterTurnover(ctx *TurnoverContext) {}

// ExitTurnover is called when production turnover is exited.
func (s *BaseTurnoverListener) ExitTurnover(ctx *TurnoverContext) {}
