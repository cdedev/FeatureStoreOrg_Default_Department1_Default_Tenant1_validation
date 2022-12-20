// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671516081858/OccupationCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OccupationCode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOccupationCodeListener is a complete listener for a parse tree produced by OccupationCodeParser.
type BaseOccupationCodeListener struct{}

var _ OccupationCodeListener = &BaseOccupationCodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOccupationCodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOccupationCodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOccupationCodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOccupationCodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOccupationCodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOccupationCodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOccupationcode is called when production occupationcode is entered.
func (s *BaseOccupationCodeListener) EnterOccupationcode(ctx *OccupationcodeContext) {}

// ExitOccupationcode is called when production occupationcode is exited.
func (s *BaseOccupationCodeListener) ExitOccupationcode(ctx *OccupationcodeContext) {}
