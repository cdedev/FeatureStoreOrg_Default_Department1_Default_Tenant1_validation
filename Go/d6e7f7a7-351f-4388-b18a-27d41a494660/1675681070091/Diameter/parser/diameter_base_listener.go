// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675681070091/Diameter.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diameter

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDiameterListener is a complete listener for a parse tree produced by DiameterParser.
type BaseDiameterListener struct{}

var _ DiameterListener = &BaseDiameterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDiameterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDiameterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDiameterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDiameterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDiameterListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDiameterListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDiameter is called when production diameter is entered.
func (s *BaseDiameterListener) EnterDiameter(ctx *DiameterContext) {}

// ExitDiameter is called when production diameter is exited.
func (s *BaseDiameterListener) ExitDiameter(ctx *DiameterContext) {}
