// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669629685893/Churn.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Churn

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChurnListener is a complete listener for a parse tree produced by ChurnParser.
type BaseChurnListener struct{}

var _ ChurnListener = &BaseChurnListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChurnListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChurnListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChurnListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChurnListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseChurnListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseChurnListener) ExitExpression(ctx *ExpressionContext) {}

// EnterChurn is called when production churn is entered.
func (s *BaseChurnListener) EnterChurn(ctx *ChurnContext) {}

// ExitChurn is called when production churn is exited.
func (s *BaseChurnListener) ExitChurn(ctx *ChurnContext) {}
