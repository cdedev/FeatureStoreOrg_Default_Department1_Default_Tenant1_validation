// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675745067273/HeadQuarters.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // HeadQuarters

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHeadQuartersListener is a complete listener for a parse tree produced by HeadQuartersParser.
type BaseHeadQuartersListener struct{}

var _ HeadQuartersListener = &BaseHeadQuartersListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHeadQuartersListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHeadQuartersListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHeadQuartersListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHeadQuartersListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHeadQuartersListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHeadQuartersListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHeadquarters is called when production headquarters is entered.
func (s *BaseHeadQuartersListener) EnterHeadquarters(ctx *HeadquartersContext) {}

// ExitHeadquarters is called when production headquarters is exited.
func (s *BaseHeadQuartersListener) ExitHeadquarters(ctx *HeadquartersContext) {}
