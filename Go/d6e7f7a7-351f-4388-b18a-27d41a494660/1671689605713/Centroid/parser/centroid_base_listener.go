// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671689605713/Centroid.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Centroid

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCentroidListener is a complete listener for a parse tree produced by CentroidParser.
type BaseCentroidListener struct{}

var _ CentroidListener = &BaseCentroidListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCentroidListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCentroidListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCentroidListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCentroidListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCentroidListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCentroidListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCentroid is called when production centroid is entered.
func (s *BaseCentroidListener) EnterCentroid(ctx *CentroidContext) {}

// ExitCentroid is called when production centroid is exited.
func (s *BaseCentroidListener) ExitCentroid(ctx *CentroidContext) {}
