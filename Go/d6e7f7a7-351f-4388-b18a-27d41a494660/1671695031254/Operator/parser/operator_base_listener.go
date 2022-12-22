// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671695031254/Operator.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Operator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOperatorListener is a complete listener for a parse tree produced by OperatorParser.
type BaseOperatorListener struct{}

var _ OperatorListener = &BaseOperatorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOperatorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOperatorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOperatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOperatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOperatorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOperatorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseOperatorListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseOperatorListener) ExitOperator(ctx *OperatorContext) {}
