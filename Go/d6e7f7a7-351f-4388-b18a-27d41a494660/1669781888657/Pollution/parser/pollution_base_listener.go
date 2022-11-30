// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669781888657/Pollution.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pollution

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePollutionListener is a complete listener for a parse tree produced by PollutionParser.
type BasePollutionListener struct{}

var _ PollutionListener = &BasePollutionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePollutionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePollutionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePollutionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePollutionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePollutionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePollutionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPollution is called when production pollution is entered.
func (s *BasePollutionListener) EnterPollution(ctx *PollutionContext) {}

// ExitPollution is called when production pollution is exited.
func (s *BasePollutionListener) ExitPollution(ctx *PollutionContext) {}
