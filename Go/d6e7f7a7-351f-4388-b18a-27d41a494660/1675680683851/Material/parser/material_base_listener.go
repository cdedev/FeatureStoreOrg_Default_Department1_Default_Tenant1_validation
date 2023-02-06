// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675680683851/Material.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Material

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMaterialListener is a complete listener for a parse tree produced by MaterialParser.
type BaseMaterialListener struct{}

var _ MaterialListener = &BaseMaterialListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMaterialListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMaterialListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMaterialListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMaterialListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMaterialListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMaterialListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMaterial is called when production material is entered.
func (s *BaseMaterialListener) EnterMaterial(ctx *MaterialContext) {}

// ExitMaterial is called when production material is exited.
func (s *BaseMaterialListener) ExitMaterial(ctx *MaterialContext) {}
