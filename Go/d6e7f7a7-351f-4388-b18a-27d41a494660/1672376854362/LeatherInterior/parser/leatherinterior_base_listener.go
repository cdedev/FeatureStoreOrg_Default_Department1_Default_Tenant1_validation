// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376854362/LeatherInterior.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // LeatherInterior

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLeatherInteriorListener is a complete listener for a parse tree produced by LeatherInteriorParser.
type BaseLeatherInteriorListener struct{}

var _ LeatherInteriorListener = &BaseLeatherInteriorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLeatherInteriorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLeatherInteriorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLeatherInteriorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLeatherInteriorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLeatherInteriorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLeatherInteriorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLeatherinterior is called when production leatherinterior is entered.
func (s *BaseLeatherInteriorListener) EnterLeatherinterior(ctx *LeatherinteriorContext) {}

// ExitLeatherinterior is called when production leatherinterior is exited.
func (s *BaseLeatherInteriorListener) ExitLeatherinterior(ctx *LeatherinteriorContext) {}
