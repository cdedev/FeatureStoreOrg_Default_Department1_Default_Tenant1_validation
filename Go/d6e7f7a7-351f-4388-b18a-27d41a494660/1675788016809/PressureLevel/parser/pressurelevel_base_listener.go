// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675788016809/PressureLevel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PressureLevel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePressureLevelListener is a complete listener for a parse tree produced by PressureLevelParser.
type BasePressureLevelListener struct{}

var _ PressureLevelListener = &BasePressureLevelListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePressureLevelListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePressureLevelListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePressureLevelListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePressureLevelListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePressureLevelListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePressureLevelListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPressurelevel is called when production pressurelevel is entered.
func (s *BasePressureLevelListener) EnterPressurelevel(ctx *PressurelevelContext) {}

// ExitPressurelevel is called when production pressurelevel is exited.
func (s *BasePressureLevelListener) ExitPressurelevel(ctx *PressurelevelContext) {}
