// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675845518714/Constituency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Constituency

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConstituencyListener is a complete listener for a parse tree produced by ConstituencyParser.
type BaseConstituencyListener struct{}

var _ ConstituencyListener = &BaseConstituencyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConstituencyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConstituencyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConstituencyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConstituencyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConstituencyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConstituencyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConstituency is called when production constituency is entered.
func (s *BaseConstituencyListener) EnterConstituency(ctx *ConstituencyContext) {}

// ExitConstituency is called when production constituency is exited.
func (s *BaseConstituencyListener) ExitConstituency(ctx *ConstituencyContext) {}
