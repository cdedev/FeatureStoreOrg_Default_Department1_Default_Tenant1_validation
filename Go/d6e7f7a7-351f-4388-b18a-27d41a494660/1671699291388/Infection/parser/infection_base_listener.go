// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671699291388/Infection.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Infection

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInfectionListener is a complete listener for a parse tree produced by InfectionParser.
type BaseInfectionListener struct{}

var _ InfectionListener = &BaseInfectionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInfectionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInfectionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInfectionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInfectionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInfectionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInfectionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInfection is called when production infection is entered.
func (s *BaseInfectionListener) EnterInfection(ctx *InfectionContext) {}

// ExitInfection is called when production infection is exited.
func (s *BaseInfectionListener) ExitInfection(ctx *InfectionContext) {}
