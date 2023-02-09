// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675931105337/Headache.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Headache

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHeadacheListener is a complete listener for a parse tree produced by HeadacheParser.
type BaseHeadacheListener struct{}

var _ HeadacheListener = &BaseHeadacheListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHeadacheListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHeadacheListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHeadacheListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHeadacheListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHeadacheListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHeadacheListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHeadache is called when production headache is entered.
func (s *BaseHeadacheListener) EnterHeadache(ctx *HeadacheContext) {}

// ExitHeadache is called when production headache is exited.
func (s *BaseHeadacheListener) ExitHeadache(ctx *HeadacheContext) {}
