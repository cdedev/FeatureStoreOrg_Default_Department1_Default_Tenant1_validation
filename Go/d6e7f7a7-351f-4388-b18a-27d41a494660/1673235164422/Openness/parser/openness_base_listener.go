// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235164422/Openness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Openness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOpennessListener is a complete listener for a parse tree produced by OpennessParser.
type BaseOpennessListener struct{}

var _ OpennessListener = &BaseOpennessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOpennessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOpennessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOpennessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOpennessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOpennessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOpennessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOpenness is called when production openness is entered.
func (s *BaseOpennessListener) EnterOpenness(ctx *OpennessContext) {}

// ExitOpenness is called when production openness is exited.
func (s *BaseOpennessListener) ExitOpenness(ctx *OpennessContext) {}
