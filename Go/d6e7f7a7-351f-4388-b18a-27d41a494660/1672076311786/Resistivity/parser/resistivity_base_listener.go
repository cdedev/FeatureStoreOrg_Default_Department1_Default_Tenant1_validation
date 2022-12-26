// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076311786/Resistivity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Resistivity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseResistivityListener is a complete listener for a parse tree produced by ResistivityParser.
type BaseResistivityListener struct{}

var _ ResistivityListener = &BaseResistivityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseResistivityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseResistivityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseResistivityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseResistivityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseResistivityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseResistivityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterResistivity is called when production resistivity is entered.
func (s *BaseResistivityListener) EnterResistivity(ctx *ResistivityContext) {}

// ExitResistivity is called when production resistivity is exited.
func (s *BaseResistivityListener) ExitResistivity(ctx *ResistivityContext) {}
