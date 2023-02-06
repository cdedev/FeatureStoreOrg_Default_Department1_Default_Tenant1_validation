// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675673022752/Correlation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Correlation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCorrelationListener is a complete listener for a parse tree produced by CorrelationParser.
type BaseCorrelationListener struct{}

var _ CorrelationListener = &BaseCorrelationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCorrelationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCorrelationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCorrelationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCorrelationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCorrelationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCorrelationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCorrelation is called when production correlation is entered.
func (s *BaseCorrelationListener) EnterCorrelation(ctx *CorrelationContext) {}

// ExitCorrelation is called when production correlation is exited.
func (s *BaseCorrelationListener) ExitCorrelation(ctx *CorrelationContext) {}
