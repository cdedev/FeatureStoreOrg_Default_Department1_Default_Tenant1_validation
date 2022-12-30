// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377318029/DualSim.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DualSim

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDualSimListener is a complete listener for a parse tree produced by DualSimParser.
type BaseDualSimListener struct{}

var _ DualSimListener = &BaseDualSimListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDualSimListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDualSimListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDualSimListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDualSimListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDualSimListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDualSimListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDualsim is called when production dualsim is entered.
func (s *BaseDualSimListener) EnterDualsim(ctx *DualsimContext) {}

// ExitDualsim is called when production dualsim is exited.
func (s *BaseDualSimListener) ExitDualsim(ctx *DualsimContext) {}
