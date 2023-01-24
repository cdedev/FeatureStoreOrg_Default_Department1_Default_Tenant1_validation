// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674533761260/Urban.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Urban

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseUrbanListener is a complete listener for a parse tree produced by UrbanParser.
type BaseUrbanListener struct{}

var _ UrbanListener = &BaseUrbanListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseUrbanListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseUrbanListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseUrbanListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseUrbanListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseUrbanListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseUrbanListener) ExitExpression(ctx *ExpressionContext) {}

// EnterUrban is called when production urban is entered.
func (s *BaseUrbanListener) EnterUrban(ctx *UrbanContext) {}

// ExitUrban is called when production urban is exited.
func (s *BaseUrbanListener) ExitUrban(ctx *UrbanContext) {}
