// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077078876/AprMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AprMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAprMaxtempListener is a complete listener for a parse tree produced by AprMaxtempParser.
type BaseAprMaxtempListener struct{}

var _ AprMaxtempListener = &BaseAprMaxtempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAprMaxtempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAprMaxtempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAprMaxtempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAprMaxtempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAprMaxtempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAprMaxtempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAprmaxtemp is called when production aprmaxtemp is entered.
func (s *BaseAprMaxtempListener) EnterAprmaxtemp(ctx *AprmaxtempContext) {}

// ExitAprmaxtemp is called when production aprmaxtemp is exited.
func (s *BaseAprMaxtempListener) ExitAprmaxtemp(ctx *AprmaxtempContext) {}
