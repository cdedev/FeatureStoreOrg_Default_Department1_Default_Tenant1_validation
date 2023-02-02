// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675307034189/Diastolic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diastolic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDiastolicListener is a complete listener for a parse tree produced by DiastolicParser.
type BaseDiastolicListener struct{}

var _ DiastolicListener = &BaseDiastolicListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDiastolicListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDiastolicListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDiastolicListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDiastolicListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDiastolicListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDiastolicListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDiastolic is called when production diastolic is entered.
func (s *BaseDiastolicListener) EnterDiastolic(ctx *DiastolicContext) {}

// ExitDiastolic is called when production diastolic is exited.
func (s *BaseDiastolicListener) ExitDiastolic(ctx *DiastolicContext) {}
