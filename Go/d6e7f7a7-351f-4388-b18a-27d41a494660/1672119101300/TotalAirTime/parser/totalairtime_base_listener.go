// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119101300/TotalAirTime.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TotalAirTime

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTotalAirTimeListener is a complete listener for a parse tree produced by TotalAirTimeParser.
type BaseTotalAirTimeListener struct{}

var _ TotalAirTimeListener = &BaseTotalAirTimeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTotalAirTimeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTotalAirTimeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTotalAirTimeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTotalAirTimeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTotalAirTimeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTotalAirTimeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTotalairtime is called when production totalairtime is entered.
func (s *BaseTotalAirTimeListener) EnterTotalairtime(ctx *TotalairtimeContext) {}

// ExitTotalairtime is called when production totalairtime is exited.
func (s *BaseTotalAirTimeListener) ExitTotalairtime(ctx *TotalairtimeContext) {}
