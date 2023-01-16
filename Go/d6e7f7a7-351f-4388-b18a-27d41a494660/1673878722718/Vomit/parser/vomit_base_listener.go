// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878722718/Vomit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vomit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVomitListener is a complete listener for a parse tree produced by VomitParser.
type BaseVomitListener struct{}

var _ VomitListener = &BaseVomitListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVomitListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVomitListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVomitListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVomitListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVomitListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVomitListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVomit is called when production vomit is entered.
func (s *BaseVomitListener) EnterVomit(ctx *VomitContext) {}

// ExitVomit is called when production vomit is exited.
func (s *BaseVomitListener) ExitVomit(ctx *VomitContext) {}
