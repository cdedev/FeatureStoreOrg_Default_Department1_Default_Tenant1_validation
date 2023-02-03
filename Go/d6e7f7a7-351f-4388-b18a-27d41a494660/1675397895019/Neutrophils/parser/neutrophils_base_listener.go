// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675397895019/Neutrophils.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Neutrophils

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNeutrophilsListener is a complete listener for a parse tree produced by NeutrophilsParser.
type BaseNeutrophilsListener struct{}

var _ NeutrophilsListener = &BaseNeutrophilsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNeutrophilsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNeutrophilsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNeutrophilsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNeutrophilsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNeutrophilsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNeutrophilsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNeutrophils is called when production neutrophils is entered.
func (s *BaseNeutrophilsListener) EnterNeutrophils(ctx *NeutrophilsContext) {}

// ExitNeutrophils is called when production neutrophils is exited.
func (s *BaseNeutrophilsListener) ExitNeutrophils(ctx *NeutrophilsContext) {}
