// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238774423/Normal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Normal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNormalListener is a complete listener for a parse tree produced by NormalParser.
type BaseNormalListener struct{}

var _ NormalListener = &BaseNormalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNormalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNormalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNormalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNormalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNormalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNormalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNormal is called when production normal is entered.
func (s *BaseNormalListener) EnterNormal(ctx *NormalContext) {}

// ExitNormal is called when production normal is exited.
func (s *BaseNormalListener) ExitNormal(ctx *NormalContext) {}
