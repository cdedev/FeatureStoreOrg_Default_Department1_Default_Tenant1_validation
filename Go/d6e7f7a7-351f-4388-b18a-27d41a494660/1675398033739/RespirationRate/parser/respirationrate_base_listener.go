// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675398033739/RespirationRate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RespirationRate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRespirationRateListener is a complete listener for a parse tree produced by RespirationRateParser.
type BaseRespirationRateListener struct{}

var _ RespirationRateListener = &BaseRespirationRateListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRespirationRateListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRespirationRateListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRespirationRateListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRespirationRateListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRespirationRateListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRespirationRateListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRespirationrate is called when production respirationrate is entered.
func (s *BaseRespirationRateListener) EnterRespirationrate(ctx *RespirationrateContext) {}

// ExitRespirationrate is called when production respirationrate is exited.
func (s *BaseRespirationRateListener) ExitRespirationrate(ctx *RespirationrateContext) {}
