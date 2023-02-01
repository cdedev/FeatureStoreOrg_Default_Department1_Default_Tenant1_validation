// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675231501618/HoursWorked.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // HoursWorked

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHoursWorkedListener is a complete listener for a parse tree produced by HoursWorkedParser.
type BaseHoursWorkedListener struct{}

var _ HoursWorkedListener = &BaseHoursWorkedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHoursWorkedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHoursWorkedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHoursWorkedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHoursWorkedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHoursWorkedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHoursWorkedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHoursworked is called when production hoursworked is entered.
func (s *BaseHoursWorkedListener) EnterHoursworked(ctx *HoursworkedContext) {}

// ExitHoursworked is called when production hoursworked is exited.
func (s *BaseHoursWorkedListener) ExitHoursworked(ctx *HoursworkedContext) {}
