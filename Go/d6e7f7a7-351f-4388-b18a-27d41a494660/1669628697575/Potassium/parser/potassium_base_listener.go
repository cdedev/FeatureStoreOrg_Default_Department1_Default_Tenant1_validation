// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628697575/Potassium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Potassium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePotassiumListener is a complete listener for a parse tree produced by PotassiumParser.
type BasePotassiumListener struct{}

var _ PotassiumListener = &BasePotassiumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePotassiumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePotassiumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePotassiumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePotassiumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePotassiumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePotassiumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPotassium is called when production potassium is entered.
func (s *BasePotassiumListener) EnterPotassium(ctx *PotassiumContext) {}

// ExitPotassium is called when production potassium is exited.
func (s *BasePotassiumListener) ExitPotassium(ctx *PotassiumContext) {}
