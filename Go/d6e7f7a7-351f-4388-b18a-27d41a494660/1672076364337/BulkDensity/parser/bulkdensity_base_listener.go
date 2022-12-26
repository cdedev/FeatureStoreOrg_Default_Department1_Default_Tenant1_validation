// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076364337/BulkDensity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BulkDensity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBulkDensityListener is a complete listener for a parse tree produced by BulkDensityParser.
type BaseBulkDensityListener struct{}

var _ BulkDensityListener = &BaseBulkDensityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBulkDensityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBulkDensityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBulkDensityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBulkDensityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBulkDensityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBulkDensityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBulkdensity is called when production bulkdensity is entered.
func (s *BaseBulkDensityListener) EnterBulkdensity(ctx *BulkdensityContext) {}

// ExitBulkdensity is called when production bulkdensity is exited.
func (s *BaseBulkDensityListener) ExitBulkdensity(ctx *BulkdensityContext) {}
