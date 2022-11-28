// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669631036575/Drug.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Drug

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDrugListener is a complete listener for a parse tree produced by DrugParser.
type BaseDrugListener struct{}

var _ DrugListener = &BaseDrugListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDrugListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDrugListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDrugListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDrugListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDrugListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDrugListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDrug is called when production drug is entered.
func (s *BaseDrugListener) EnterDrug(ctx *DrugContext) {}

// ExitDrug is called when production drug is exited.
func (s *BaseDrugListener) ExitDrug(ctx *DrugContext) {}
