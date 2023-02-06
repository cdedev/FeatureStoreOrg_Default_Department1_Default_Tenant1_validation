// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675674332598/Disease.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Disease

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDiseaseListener is a complete listener for a parse tree produced by DiseaseParser.
type BaseDiseaseListener struct{}

var _ DiseaseListener = &BaseDiseaseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDiseaseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDiseaseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDiseaseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDiseaseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDiseaseListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDiseaseListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDisease is called when production disease is entered.
func (s *BaseDiseaseListener) EnterDisease(ctx *DiseaseContext) {}

// ExitDisease is called when production disease is exited.
func (s *BaseDiseaseListener) ExitDisease(ctx *DiseaseContext) {}
