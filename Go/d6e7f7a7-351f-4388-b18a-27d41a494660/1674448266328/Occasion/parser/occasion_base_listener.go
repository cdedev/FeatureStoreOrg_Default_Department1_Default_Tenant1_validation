// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674448266328/Occasion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Occasion

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOccasionListener is a complete listener for a parse tree produced by OccasionParser.
type BaseOccasionListener struct{}

var _ OccasionListener = &BaseOccasionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOccasionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOccasionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOccasionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOccasionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOccasionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOccasionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOccasion is called when production occasion is entered.
func (s *BaseOccasionListener) EnterOccasion(ctx *OccasionContext) {}

// ExitOccasion is called when production occasion is exited.
func (s *BaseOccasionListener) ExitOccasion(ctx *OccasionContext) {}
