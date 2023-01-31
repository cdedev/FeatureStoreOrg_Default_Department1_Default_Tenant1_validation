// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675151690590/Acres.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Acres

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAcresListener is a complete listener for a parse tree produced by AcresParser.
type BaseAcresListener struct{}

var _ AcresListener = &BaseAcresListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAcresListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAcresListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAcresListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAcresListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAcresListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAcresListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAcres is called when production acres is entered.
func (s *BaseAcresListener) EnterAcres(ctx *AcresContext) {}

// ExitAcres is called when production acres is exited.
func (s *BaseAcresListener) ExitAcres(ctx *AcresContext) {}
