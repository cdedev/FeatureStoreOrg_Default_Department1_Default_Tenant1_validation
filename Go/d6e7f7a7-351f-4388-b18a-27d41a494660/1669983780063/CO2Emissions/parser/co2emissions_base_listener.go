// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669983780063/CO2Emissions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CO2Emissions

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCO2EmissionsListener is a complete listener for a parse tree produced by CO2EmissionsParser.
type BaseCO2EmissionsListener struct{}

var _ CO2EmissionsListener = &BaseCO2EmissionsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCO2EmissionsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCO2EmissionsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCO2EmissionsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCO2EmissionsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCO2EmissionsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCO2EmissionsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCo2emissions is called when production co2emissions is entered.
func (s *BaseCO2EmissionsListener) EnterCo2emissions(ctx *Co2emissionsContext) {}

// ExitCo2emissions is called when production co2emissions is exited.
func (s *BaseCO2EmissionsListener) ExitCo2emissions(ctx *Co2emissionsContext) {}
