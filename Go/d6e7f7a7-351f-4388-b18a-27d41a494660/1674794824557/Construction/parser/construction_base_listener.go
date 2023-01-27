// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674794824557/Construction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Construction

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConstructionListener is a complete listener for a parse tree produced by ConstructionParser.
type BaseConstructionListener struct{}

var _ ConstructionListener = &BaseConstructionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConstructionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConstructionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConstructionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConstructionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConstructionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConstructionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConstruction is called when production construction is entered.
func (s *BaseConstructionListener) EnterConstruction(ctx *ConstructionContext) {}

// ExitConstruction is called when production construction is exited.
func (s *BaseConstructionListener) ExitConstruction(ctx *ConstructionContext) {}
