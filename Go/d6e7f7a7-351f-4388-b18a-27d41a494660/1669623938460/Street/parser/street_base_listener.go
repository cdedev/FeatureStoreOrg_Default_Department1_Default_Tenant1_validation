// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623938460/Street.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Street

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStreetListener is a complete listener for a parse tree produced by StreetParser.
type BaseStreetListener struct{}

var _ StreetListener = &BaseStreetListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStreetListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStreetListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStreetListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStreetListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStreetListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStreetListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStreet is called when production street is entered.
func (s *BaseStreetListener) EnterStreet(ctx *StreetContext) {}

// ExitStreet is called when production street is exited.
func (s *BaseStreetListener) ExitStreet(ctx *StreetContext) {}
