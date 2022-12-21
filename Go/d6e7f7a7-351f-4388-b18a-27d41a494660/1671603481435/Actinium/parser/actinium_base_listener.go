// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603481435/Actinium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Actinium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseActiniumListener is a complete listener for a parse tree produced by ActiniumParser.
type BaseActiniumListener struct{}

var _ ActiniumListener = &BaseActiniumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseActiniumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseActiniumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseActiniumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseActiniumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseActiniumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseActiniumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterActinium is called when production actinium is entered.
func (s *BaseActiniumListener) EnterActinium(ctx *ActiniumContext) {}

// ExitActinium is called when production actinium is exited.
func (s *BaseActiniumListener) ExitActinium(ctx *ActiniumContext) {}
