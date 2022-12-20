// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671513436611/Quarantine.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Quarantine

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQuarantineListener is a complete listener for a parse tree produced by QuarantineParser.
type BaseQuarantineListener struct{}

var _ QuarantineListener = &BaseQuarantineListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQuarantineListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQuarantineListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQuarantineListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQuarantineListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseQuarantineListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseQuarantineListener) ExitExpression(ctx *ExpressionContext) {}

// EnterQuarantine is called when production quarantine is entered.
func (s *BaseQuarantineListener) EnterQuarantine(ctx *QuarantineContext) {}

// ExitQuarantine is called when production quarantine is exited.
func (s *BaseQuarantineListener) ExitQuarantine(ctx *QuarantineContext) {}
