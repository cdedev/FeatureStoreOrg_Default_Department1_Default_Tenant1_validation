// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675227584189/District.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // District

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDistrictListener is a complete listener for a parse tree produced by DistrictParser.
type BaseDistrictListener struct{}

var _ DistrictListener = &BaseDistrictListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDistrictListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDistrictListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDistrictListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDistrictListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDistrictListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDistrictListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDistrict is called when production district is entered.
func (s *BaseDistrictListener) EnterDistrict(ctx *DistrictContext) {}

// ExitDistrict is called when production district is exited.
func (s *BaseDistrictListener) ExitDistrict(ctx *DistrictContext) {}
