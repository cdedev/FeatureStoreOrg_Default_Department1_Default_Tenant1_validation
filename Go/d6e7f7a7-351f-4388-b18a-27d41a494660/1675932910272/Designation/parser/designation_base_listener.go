// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675932910272/Designation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Designation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDesignationListener is a complete listener for a parse tree produced by DesignationParser.
type BaseDesignationListener struct{}

var _ DesignationListener = &BaseDesignationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDesignationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDesignationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDesignationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDesignationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDesignationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDesignationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDesignation is called when production designation is entered.
func (s *BaseDesignationListener) EnterDesignation(ctx *DesignationContext) {}

// ExitDesignation is called when production designation is exited.
func (s *BaseDesignationListener) ExitDesignation(ctx *DesignationContext) {}
