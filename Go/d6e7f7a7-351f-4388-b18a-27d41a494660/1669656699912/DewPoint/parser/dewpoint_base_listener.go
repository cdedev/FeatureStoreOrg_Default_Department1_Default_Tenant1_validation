// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656699912/DewPoint.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DewPoint

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDewPointListener is a complete listener for a parse tree produced by DewPointParser.
type BaseDewPointListener struct{}

var _ DewPointListener = &BaseDewPointListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDewPointListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDewPointListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDewPointListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDewPointListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDewPointListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDewPointListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDewpoint is called when production dewpoint is entered.
func (s *BaseDewPointListener) EnterDewpoint(ctx *DewpointContext) {}

// ExitDewpoint is called when production dewpoint is exited.
func (s *BaseDewPointListener) ExitDewpoint(ctx *DewpointContext) {}
