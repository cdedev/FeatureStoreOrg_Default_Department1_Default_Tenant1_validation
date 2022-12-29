// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284768398/CartoonReference.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CartoonReference

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCartoonReferenceListener is a complete listener for a parse tree produced by CartoonReferenceParser.
type BaseCartoonReferenceListener struct{}

var _ CartoonReferenceListener = &BaseCartoonReferenceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCartoonReferenceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCartoonReferenceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCartoonReferenceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCartoonReferenceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCartoonReferenceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCartoonReferenceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCartoonreference is called when production cartoonreference is entered.
func (s *BaseCartoonReferenceListener) EnterCartoonreference(ctx *CartoonreferenceContext) {}

// ExitCartoonreference is called when production cartoonreference is exited.
func (s *BaseCartoonReferenceListener) ExitCartoonreference(ctx *CartoonreferenceContext) {}
