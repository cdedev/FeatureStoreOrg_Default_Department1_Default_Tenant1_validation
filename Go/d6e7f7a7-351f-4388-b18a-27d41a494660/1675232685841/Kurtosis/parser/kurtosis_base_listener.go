// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675232685841/Kurtosis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Kurtosis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseKurtosisListener is a complete listener for a parse tree produced by KurtosisParser.
type BaseKurtosisListener struct{}

var _ KurtosisListener = &BaseKurtosisListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKurtosisListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKurtosisListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKurtosisListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKurtosisListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseKurtosisListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseKurtosisListener) ExitExpression(ctx *ExpressionContext) {}

// EnterKurtosis is called when production kurtosis is entered.
func (s *BaseKurtosisListener) EnterKurtosis(ctx *KurtosisContext) {}

// ExitKurtosis is called when production kurtosis is exited.
func (s *BaseKurtosisListener) ExitKurtosis(ctx *KurtosisContext) {}
