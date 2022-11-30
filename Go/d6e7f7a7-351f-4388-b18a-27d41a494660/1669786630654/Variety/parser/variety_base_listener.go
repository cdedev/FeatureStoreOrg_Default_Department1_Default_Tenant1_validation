// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669786630654/Variety.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Variety

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVarietyListener is a complete listener for a parse tree produced by VarietyParser.
type BaseVarietyListener struct{}

var _ VarietyListener = &BaseVarietyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVarietyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVarietyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVarietyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVarietyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVarietyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVarietyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVariety is called when production variety is entered.
func (s *BaseVarietyListener) EnterVariety(ctx *VarietyContext) {}

// ExitVariety is called when production variety is exited.
func (s *BaseVarietyListener) ExitVariety(ctx *VarietyContext) {}
