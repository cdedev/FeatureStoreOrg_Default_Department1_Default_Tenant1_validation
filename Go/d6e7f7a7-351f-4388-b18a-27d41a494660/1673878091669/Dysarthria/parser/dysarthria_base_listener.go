// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878091669/Dysarthria.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dysarthria

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDysarthriaListener is a complete listener for a parse tree produced by DysarthriaParser.
type BaseDysarthriaListener struct{}

var _ DysarthriaListener = &BaseDysarthriaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDysarthriaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDysarthriaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDysarthriaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDysarthriaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDysarthriaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDysarthriaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDysarthria is called when production dysarthria is entered.
func (s *BaseDysarthriaListener) EnterDysarthria(ctx *DysarthriaContext) {}

// ExitDysarthria is called when production dysarthria is exited.
func (s *BaseDysarthriaListener) ExitDysarthria(ctx *DysarthriaContext) {}
