// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674532543597/Electors.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Electors

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseElectorsListener is a complete listener for a parse tree produced by ElectorsParser.
type BaseElectorsListener struct{}

var _ ElectorsListener = &BaseElectorsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseElectorsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseElectorsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseElectorsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseElectorsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseElectorsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseElectorsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterElectors is called when production electors is entered.
func (s *BaseElectorsListener) EnterElectors(ctx *ElectorsContext) {}

// ExitElectors is called when production electors is exited.
func (s *BaseElectorsListener) ExitElectors(ctx *ElectorsContext) {}
