// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072682011/C02Emissions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // C02Emissions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseC02EmissionsListener is a complete listener for a parse tree produced by C02EmissionsParser.
type BaseC02EmissionsListener struct{}

var _ C02EmissionsListener = &BaseC02EmissionsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseC02EmissionsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseC02EmissionsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseC02EmissionsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseC02EmissionsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseC02EmissionsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseC02EmissionsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterC02emissions is called when production c02emissions is entered.
func (s *BaseC02EmissionsListener) EnterC02emissions(ctx *C02emissionsContext) {}

// ExitC02emissions is called when production c02emissions is exited.
func (s *BaseC02EmissionsListener) ExitC02emissions(ctx *C02emissionsContext) {}
