// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235389726/Extraversion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Extraversion

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExtraversionListener is a complete listener for a parse tree produced by ExtraversionParser.
type BaseExtraversionListener struct{}

var _ ExtraversionListener = &BaseExtraversionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExtraversionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExtraversionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExtraversionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExtraversionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExtraversionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExtraversionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExtraversion is called when production extraversion is entered.
func (s *BaseExtraversionListener) EnterExtraversion(ctx *ExtraversionContext) {}

// ExitExtraversion is called when production extraversion is exited.
func (s *BaseExtraversionListener) ExitExtraversion(ctx *ExtraversionContext) {}
