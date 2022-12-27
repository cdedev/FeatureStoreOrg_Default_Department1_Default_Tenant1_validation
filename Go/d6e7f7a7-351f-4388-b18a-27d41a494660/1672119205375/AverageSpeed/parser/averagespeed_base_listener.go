// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119205375/AverageSpeed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AverageSpeed

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAverageSpeedListener is a complete listener for a parse tree produced by AverageSpeedParser.
type BaseAverageSpeedListener struct{}

var _ AverageSpeedListener = &BaseAverageSpeedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAverageSpeedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAverageSpeedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAverageSpeedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAverageSpeedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAverageSpeedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAverageSpeedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAveragespeed is called when production averagespeed is entered.
func (s *BaseAverageSpeedListener) EnterAveragespeed(ctx *AveragespeedContext) {}

// ExitAveragespeed is called when production averagespeed is exited.
func (s *BaseAverageSpeedListener) ExitAveragespeed(ctx *AveragespeedContext) {}
