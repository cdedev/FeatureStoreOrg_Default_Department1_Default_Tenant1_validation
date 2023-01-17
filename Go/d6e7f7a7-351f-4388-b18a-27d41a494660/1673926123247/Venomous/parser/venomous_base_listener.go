// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673926123247/Venomous.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Venomous

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVenomousListener is a complete listener for a parse tree produced by VenomousParser.
type BaseVenomousListener struct{}

var _ VenomousListener = &BaseVenomousListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVenomousListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVenomousListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVenomousListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVenomousListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVenomousListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVenomousListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVenomous is called when production venomous is entered.
func (s *BaseVenomousListener) EnterVenomous(ctx *VenomousContext) {}

// ExitVenomous is called when production venomous is exited.
func (s *BaseVenomousListener) ExitVenomous(ctx *VenomousContext) {}
