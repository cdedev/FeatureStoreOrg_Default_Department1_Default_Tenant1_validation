// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669983891898/Cylinders.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cylinders

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCylindersListener is a complete listener for a parse tree produced by CylindersParser.
type BaseCylindersListener struct{}

var _ CylindersListener = &BaseCylindersListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCylindersListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCylindersListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCylindersListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCylindersListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCylindersListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCylindersListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCylinders is called when production cylinders is entered.
func (s *BaseCylindersListener) EnterCylinders(ctx *CylindersContext) {}

// ExitCylinders is called when production cylinders is exited.
func (s *BaseCylindersListener) ExitCylinders(ctx *CylindersContext) {}
