// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674839818224/Joblost.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Joblost

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJoblostListener is a complete listener for a parse tree produced by JoblostParser.
type BaseJoblostListener struct{}

var _ JoblostListener = &BaseJoblostListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJoblostListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJoblostListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJoblostListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJoblostListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJoblostListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJoblostListener) ExitExpression(ctx *ExpressionContext) {}

// EnterJoblost is called when production joblost is entered.
func (s *BaseJoblostListener) EnterJoblost(ctx *JoblostContext) {}

// ExitJoblost is called when production joblost is exited.
func (s *BaseJoblostListener) ExitJoblost(ctx *JoblostContext) {}
