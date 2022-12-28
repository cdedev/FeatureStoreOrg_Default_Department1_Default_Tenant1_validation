// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672196147965/RawEthanol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RawEthanol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRawEthanolListener is a complete listener for a parse tree produced by RawEthanolParser.
type BaseRawEthanolListener struct{}

var _ RawEthanolListener = &BaseRawEthanolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRawEthanolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRawEthanolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRawEthanolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRawEthanolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRawEthanolListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRawEthanolListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRawethanol is called when production rawethanol is entered.
func (s *BaseRawEthanolListener) EnterRawethanol(ctx *RawethanolContext) {}

// ExitRawethanol is called when production rawethanol is exited.
func (s *BaseRawEthanolListener) ExitRawethanol(ctx *RawethanolContext) {}
