// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674014957141/Jaundice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jaundice

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJaundiceListener is a complete listener for a parse tree produced by JaundiceParser.
type BaseJaundiceListener struct{}

var _ JaundiceListener = &BaseJaundiceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJaundiceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJaundiceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJaundiceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJaundiceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJaundiceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJaundiceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterJaundice is called when production jaundice is entered.
func (s *BaseJaundiceListener) EnterJaundice(ctx *JaundiceContext) {}

// ExitJaundice is called when production jaundice is exited.
func (s *BaseJaundiceListener) ExitJaundice(ctx *JaundiceContext) {}
