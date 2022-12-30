// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377742445/NCores.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NCores

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNCoresListener is a complete listener for a parse tree produced by NCoresParser.
type BaseNCoresListener struct{}

var _ NCoresListener = &BaseNCoresListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNCoresListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNCoresListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNCoresListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNCoresListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNCoresListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNCoresListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNcores is called when production ncores is entered.
func (s *BaseNCoresListener) EnterNcores(ctx *NcoresContext) {}

// ExitNcores is called when production ncores is exited.
func (s *BaseNCoresListener) ExitNcores(ctx *NcoresContext) {}
